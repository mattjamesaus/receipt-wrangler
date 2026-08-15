package repositories

import (
	"context"
	"strings"
	"time"

	"receipt-wrangler/api/internal/commands"
	"receipt-wrangler/api/internal/fx"
	"receipt-wrangler/api/internal/models"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

const defaultBaseCurrencyCode = "AUD"

// prepareReceiptMoney applies the accounting semantics shared by every receipt
// creation/update path. A rate lookup failure is represented in stored data as
// NEEDS_REVIEW; it is not a reason to discard the receipt or its evidence.
func prepareReceiptMoney(db *gorm.DB, receipt *models.Receipt, command commands.UpsertReceiptCommand, current *models.Receipt) error {
	var group models.Group
	if err := db.Model(&models.Group{}).
		Select("id", "base_currency_code").
		Where("id = ?", receipt.GroupId).
		First(&group).Error; err != nil {
		return err
	}

	baseCurrency := strings.ToUpper(strings.TrimSpace(group.BaseCurrencyCode))
	if baseCurrency == "" {
		baseCurrency = defaultBaseCurrencyCode
	}

	documentCurrency := strings.ToUpper(strings.TrimSpace(command.DocumentCurrencyCode))
	if documentCurrency == "" && current != nil && current.DocumentCurrencyCode != "" {
		documentCurrency = current.DocumentCurrencyCode
	}
	if documentCurrency == "" {
		documentCurrency = baseCurrency
	}
	receipt.DocumentCurrencyCode = documentCurrency

	if command.DocumentAmount != nil {
		receipt.DocumentAmount = *command.DocumentAmount
	} else if current != nil && documentCurrency != baseCurrency && current.DocumentCurrencyCode == documentCurrency {
		// An older client only knows Amount. For an existing foreign receipt that
		// value is the base amount, so preserve the original evidence value.
		receipt.DocumentAmount = current.DocumentAmount
	} else {
		receipt.DocumentAmount = command.Amount
	}

	if documentCurrency == baseCurrency {
		rate := decimal.NewFromInt(1)
		estimate := receipt.DocumentAmount.Round(2)
		identity := "IDENTITY"
		fxDate := dateOnly(receipt.Date)
		receipt.Amount = estimate
		receipt.EstimatedBaseAmount = &estimate
		receipt.FxRate = &rate
		receipt.FxDate = &fxDate
		receipt.FxProvider = &identity
		receipt.FxRetrievedAt = nil
		receipt.FxStatus = models.FX_DOMESTIC
		return nil
	}

	requestedStatus := command.FxStatus
	if requestedStatus == "" && current != nil &&
		(current.FxStatus == models.FX_CONFIRMED || current.FxStatus == models.FX_NEEDS_REVIEW) {
		requestedStatus = current.FxStatus
	}

	provider := fx.NewFrankfurterProvider()
	quote, quoteErr := provider.HistoricalRate(context.Background(), documentCurrency, baseCurrency, receipt.Date)
	if quoteErr != nil {
		providerName := provider.Name()
		retrievedAt := time.Now().UTC()
		receipt.FxProvider = &providerName
		receipt.FxRetrievedAt = &retrievedAt
		receipt.FxRate = nil
		receipt.FxDate = nil
		receipt.EstimatedBaseAmount = nil

		if requestedStatus == models.FX_CONFIRMED {
			receipt.Amount = command.Amount.Round(2)
			receipt.FxStatus = models.FX_CONFIRMED
			// Preserve previously captured estimate provenance when a harmless
			// edit occurs during a provider outage.
			if sameFxSource(current, receipt) && current.EstimatedBaseAmount != nil {
				receipt.EstimatedBaseAmount = current.EstimatedBaseAmount
				receipt.FxRate = current.FxRate
				receipt.FxDate = current.FxDate
				receipt.FxProvider = current.FxProvider
				receipt.FxRetrievedAt = current.FxRetrievedAt
			}
			return nil
		}
		if sameFxSource(current, receipt) && current.EstimatedBaseAmount != nil {
			receipt.Amount = current.Amount
			receipt.EstimatedBaseAmount = current.EstimatedBaseAmount
			receipt.FxRate = current.FxRate
			receipt.FxDate = current.FxDate
			receipt.FxProvider = current.FxProvider
			receipt.FxRetrievedAt = current.FxRetrievedAt
			receipt.FxStatus = models.FX_NEEDS_REVIEW
			return nil
		}

		// Zero is deliberately safer than repeating the historical bug by
		// treating a foreign document amount as though it were base currency.
		receipt.Amount = decimal.Zero
		receipt.FxStatus = models.FX_NEEDS_REVIEW
		return nil
	}

	estimate := receipt.DocumentAmount.Mul(quote.Rate).Round(2)
	receipt.EstimatedBaseAmount = &estimate
	receipt.FxRate = &quote.Rate
	receipt.FxDate = &quote.EffectiveDate
	receipt.FxProvider = &quote.Provider
	receipt.FxRetrievedAt = &quote.RetrievedAt

	if requestedStatus == models.FX_CONFIRMED {
		receipt.Amount = command.Amount.Round(2)
		receipt.FxStatus = models.FX_CONFIRMED
	} else {
		receipt.Amount = estimate
		if requestedStatus == models.FX_NEEDS_REVIEW {
			receipt.FxStatus = models.FX_NEEDS_REVIEW
		} else {
			receipt.FxStatus = models.FX_ESTIMATED
		}
	}
	return nil
}

func sameFxSource(current *models.Receipt, next *models.Receipt) bool {
	return current != nil &&
		current.DocumentCurrencyCode == next.DocumentCurrencyCode &&
		current.DocumentAmount.Equal(next.DocumentAmount) &&
		current.Date.Format("2006-01-02") == next.Date.Format("2006-01-02")
}

func dateOnly(value time.Time) time.Time {
	return time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, time.UTC)
}
