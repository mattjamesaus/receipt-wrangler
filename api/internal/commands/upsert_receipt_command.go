package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"receipt-wrangler/api/internal/models"
	"receipt-wrangler/api/internal/structs"
	"receipt-wrangler/api/internal/utils"
	"strings"
	"time"

	"github.com/shopspring/decimal"
	"golang.org/x/text/currency"
)

type UpsertReceiptCommand struct {
	Name                 string                          `json:"name"`
	Amount               decimal.Decimal                 `json:"amount"`
	DocumentAmount       *decimal.Decimal                `json:"documentAmount"`
	DocumentCurrencyCode string                          `json:"documentCurrencyCode"`
	FxStatus             models.FxStatus                 `json:"fxStatus"`
	Date                 time.Time                       `json:"date"`
	GroupId              uint                            `json:"groupId"`
	PaidByUserID         uint                            `json:"paidByUserId"`
	Status               models.ReceiptStatus            `json:"status"`
	Categories           []UpsertCategoryCommand         `json:"categories"`
	Tags                 []UpsertTagCommand              `json:"tags"`
	Items                []UpsertItemCommand             `json:"receiptItems"`
	Comments             []UpsertCommentCommand          `json:"comments"`
	CustomFields         []UpsertCustomFieldValueCommand `json:"customFields"`
	CreatedByString      string                          `json:"createdByString"`
}

func (receipt *UpsertReceiptCommand) LoadDataFromRequest(w http.ResponseWriter, r *http.Request) error {
	bytes, err := utils.GetBodyData(w, r)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &receipt)
	if err != nil {
		return err
	}

	return nil
}

func (receipt *UpsertReceiptCommand) Validate(tokenUserId uint, isCreate bool) structs.ValidatorError {
	errors := make(map[string]string)
	vErr := structs.ValidatorError{}

	if len(receipt.Name) == 0 {
		errors["name"] = "Name is required"
	}

	if receipt.Date.IsZero() {
		errors["date"] = "Date is required"
	}

	if receipt.GroupId == 0 {
		errors["groupId"] = "Group Id is required"
	}

	if receipt.PaidByUserID == 0 {
		errors["paidByUserId"] = "Paid By User Id is required"
	}

	if receipt.Status == "" {
		errors["status"] = "Status is required"
	}

	for i, category := range receipt.Categories {
		basePath := "categories." + fmt.Sprintf("%d", i)
		categoryErrors := category.Validate()
		for key, value := range categoryErrors.Errors {
			errors[basePath+"."+key] = value
		}
	}

	for i, tag := range receipt.Tags {
		basePath := "tags." + fmt.Sprintf("%d", i)
		tagErrors := tag.Validate()
		for key, value := range tagErrors.Errors {
			errors[basePath+"."+key] = value
		}
	}

	documentAmount := receipt.Amount
	if receipt.DocumentAmount != nil {
		documentAmount = *receipt.DocumentAmount
	}

	if len(receipt.DocumentCurrencyCode) > 0 {
		receipt.DocumentCurrencyCode = strings.ToUpper(strings.TrimSpace(receipt.DocumentCurrencyCode))
		if _, err := currency.ParseISO(receipt.DocumentCurrencyCode); err != nil {
			errors["documentCurrencyCode"] = "Document Currency Code must be a valid ISO 4217 code"
		}
	}

	if receipt.FxStatus != "" && !models.IsValidFxStatus(receipt.FxStatus) {
		errors["fxStatus"] = "FX Status is invalid"
	}

	for i, item := range receipt.Items {
		basePath := "receiptItems." + fmt.Sprintf("%d", i)
		// Line items are printed-evidence values, so they validate against the
		// document amount rather than the converted reporting amount.
		itemErrors := item.Validate(documentAmount, isCreate)
		for key, value := range itemErrors.Errors {
			errors[basePath+"."+key] = value
		}
	}

	for i, comment := range receipt.Comments {
		basePath := "comments." + fmt.Sprintf("%d", i)
		commentErrors := comment.Validate(tokenUserId, isCreate)
		for key, value := range commentErrors.Errors {
			errors[basePath+"."+key] = value
		}
	}

	vErr.Errors = errors
	return vErr
}

func (receipt *UpsertReceiptCommand) ToReceipt() (models.Receipt, error) {
	var result models.Receipt
	bytes, err := json.Marshal(receipt)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
