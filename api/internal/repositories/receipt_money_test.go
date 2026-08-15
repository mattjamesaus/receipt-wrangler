package repositories

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"receipt-wrangler/api/internal/commands"
	"receipt-wrangler/api/internal/models"

	"github.com/shopspring/decimal"
)

func seedMoneyGroup(t *testing.T, currencyCode string) models.Group {
	t.Helper()
	group := models.Group{Name: "money", BaseCurrencyCode: currencyCode}
	if err := GetDB().Create(&group).Error; err != nil {
		t.Fatal(err)
	}
	return group
}

func moneyCommand(groupId uint, documentCurrency string, documentAmount string) commands.UpsertReceiptCommand {
	amount := decimal.RequireFromString(documentAmount)
	return commands.UpsertReceiptCommand{
		Name: "Receipt", Amount: amount, DocumentAmount: &amount,
		DocumentCurrencyCode: documentCurrency, Date: time.Date(2025, 6, 14, 0, 0, 0, 0, time.UTC),
		GroupId: groupId, PaidByUserID: 1, Status: models.OPEN,
	}
}

func TestPrepareReceiptMoneyDomestic(t *testing.T) {
	defer TruncateTestDb()
	group := seedMoneyGroup(t, "AUD")
	command := moneyCommand(group.ID, "AUD", "12.345")
	receipt, _ := command.ToReceipt()
	if err := prepareReceiptMoney(GetDB(), &receipt, command, nil); err != nil {
		t.Fatal(err)
	}
	if receipt.FxStatus != models.FX_DOMESTIC || !receipt.Amount.Equal(decimal.RequireFromString("12.35")) {
		t.Errorf("domestic result = %#v", receipt)
	}
	if receipt.FxRate == nil || !receipt.FxRate.Equal(decimal.NewFromInt(1)) {
		t.Errorf("domestic rate = %v", receipt.FxRate)
	}
	if !receipt.DocumentAmount.Equal(decimal.RequireFromString("12.345")) || receipt.EstimatedBaseAmount == nil || !receipt.EstimatedBaseAmount.Equal(receipt.Amount) {
		t.Errorf("domestic document/estimate = %s / %v", receipt.DocumentAmount, receipt.EstimatedBaseAmount)
	}
}

func TestPrepareReceiptMoneyForeignEstimate(t *testing.T) {
	defer TruncateTestDb()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"date":"2025-06-13","base":"USD","quote":"AUD","rate":1.542}`))
	}))
	defer server.Close()
	t.Setenv("FX_PROVIDER_BASE_URL", server.URL)

	group := seedMoneyGroup(t, "AUD")
	command := moneyCommand(group.ID, "USD", "100")
	receipt, _ := command.ToReceipt()
	if err := prepareReceiptMoney(GetDB(), &receipt, command, nil); err != nil {
		t.Fatal(err)
	}
	if receipt.FxStatus != models.FX_ESTIMATED || !receipt.Amount.Equal(decimal.RequireFromString("154.20")) {
		t.Errorf("foreign result = status %s amount %s", receipt.FxStatus, receipt.Amount)
	}
	if receipt.EstimatedBaseAmount == nil || !receipt.EstimatedBaseAmount.Equal(receipt.Amount) || receipt.FxProvider == nil {
		t.Errorf("missing estimate provenance: %#v", receipt)
	}
}

func TestPrepareReceiptMoneyRateFailureNeedsReview(t *testing.T) {
	defer TruncateTestDb()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()
	t.Setenv("FX_PROVIDER_BASE_URL", server.URL)

	group := seedMoneyGroup(t, "AUD")
	command := moneyCommand(group.ID, "USD", "100")
	receipt, _ := command.ToReceipt()
	if err := prepareReceiptMoney(GetDB(), &receipt, command, nil); err != nil {
		t.Fatal(err)
	}
	if receipt.FxStatus != models.FX_NEEDS_REVIEW || !receipt.Amount.IsZero() || receipt.EstimatedBaseAmount != nil {
		t.Errorf("failure result = status %s amount %s estimate %v", receipt.FxStatus, receipt.Amount, receipt.EstimatedBaseAmount)
	}
	if receipt.DocumentCurrencyCode != "USD" || !receipt.DocumentAmount.Equal(decimal.NewFromInt(100)) {
		t.Errorf("source evidence money was not preserved: %s %s", receipt.DocumentCurrencyCode, receipt.DocumentAmount)
	}
}

func TestPrepareReceiptMoneyConfirmedKeepsEffectiveAmount(t *testing.T) {
	defer TruncateTestDb()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"date":"2025-06-13","base":"USD","quote":"AUD","rate":1.542}`))
	}))
	defer server.Close()
	t.Setenv("FX_PROVIDER_BASE_URL", server.URL)

	group := seedMoneyGroup(t, "AUD")
	command := moneyCommand(group.ID, "USD", "100")
	command.Amount = decimal.RequireFromString("153.87")
	command.FxStatus = models.FX_CONFIRMED
	receipt, _ := command.ToReceipt()
	if err := prepareReceiptMoney(GetDB(), &receipt, command, nil); err != nil {
		t.Fatal(err)
	}
	if receipt.FxStatus != models.FX_CONFIRMED || !receipt.Amount.Equal(decimal.RequireFromString("153.87")) {
		t.Errorf("confirmed result = status %s amount %s", receipt.FxStatus, receipt.Amount)
	}
	if receipt.EstimatedBaseAmount == nil || !receipt.EstimatedBaseAmount.Equal(decimal.RequireFromString("154.20")) {
		t.Errorf("estimate = %v", receipt.EstimatedBaseAmount)
	}
}

func TestPrepareReceiptMoneyKeepsExplicitNeedsReview(t *testing.T) {
	defer TruncateTestDb()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"date":"2025-06-13","base":"USD","quote":"AUD","rate":1.542}`))
	}))
	defer server.Close()
	t.Setenv("FX_PROVIDER_BASE_URL", server.URL)

	group := seedMoneyGroup(t, "AUD")
	command := moneyCommand(group.ID, "USD", "100")
	command.FxStatus = models.FX_NEEDS_REVIEW
	receipt, _ := command.ToReceipt()
	if err := prepareReceiptMoney(GetDB(), &receipt, command, nil); err != nil {
		t.Fatal(err)
	}
	if receipt.FxStatus != models.FX_NEEDS_REVIEW || !receipt.Amount.Equal(decimal.RequireFromString("154.20")) {
		t.Errorf("needs-review result = status %s amount %s", receipt.FxStatus, receipt.Amount)
	}
	if receipt.EstimatedBaseAmount == nil || receipt.FxRate == nil {
		t.Errorf("expected estimate provenance for review: %#v", receipt)
	}
}

func TestPrepareReceiptMoneyOutagePreservesExistingEstimate(t *testing.T) {
	defer TruncateTestDb()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()
	t.Setenv("FX_PROVIDER_BASE_URL", server.URL)

	group := seedMoneyGroup(t, "AUD")
	command := moneyCommand(group.ID, "USD", "100")
	priorEstimate := decimal.RequireFromString("154.20")
	priorRate := decimal.RequireFromString("1.542")
	priorDate := time.Date(2025, 6, 13, 0, 0, 0, 0, time.UTC)
	priorProvider := "Frankfurter:ECB"
	priorRetrievedAt := time.Date(2025, 6, 14, 1, 0, 0, 0, time.UTC)
	current := models.Receipt{
		Amount: priorEstimate, DocumentAmount: *command.DocumentAmount, DocumentCurrencyCode: "USD",
		EstimatedBaseAmount: &priorEstimate, FxRate: &priorRate, FxDate: &priorDate,
		FxProvider: &priorProvider, FxRetrievedAt: &priorRetrievedAt, FxStatus: models.FX_ESTIMATED,
		Date: command.Date, GroupId: group.ID,
	}
	receipt, _ := command.ToReceipt()

	if err := prepareReceiptMoney(GetDB(), &receipt, command, &current); err != nil {
		t.Fatal(err)
	}
	if receipt.FxStatus != models.FX_NEEDS_REVIEW || !receipt.Amount.Equal(priorEstimate) || receipt.FxRate == nil || !receipt.FxRate.Equal(priorRate) {
		t.Errorf("preserved outage result = status %s amount %s rate %v", receipt.FxStatus, receipt.Amount, receipt.FxRate)
	}
}
