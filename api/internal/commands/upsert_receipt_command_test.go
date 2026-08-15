package commands

import (
	"github.com/shopspring/decimal"
	"receipt-wrangler/api/internal/models"
	"receipt-wrangler/api/internal/utils"
	"testing"
	"time"
)

func validReceiptCommand() UpsertReceiptCommand {
	return UpsertReceiptCommand{
		Name:         "Test Receipt",
		Amount:       decimal.NewFromFloat(25.50),
		Date:         time.Now(),
		GroupId:      1,
		PaidByUserID: 1,
		Status:       models.ReceiptStatus("OPEN"),
	}
}

func TestUpsertReceiptCommand_Validate_ValidInputs(t *testing.T) {
	userId := uint(1)

	tests := map[string]struct {
		command  UpsertReceiptCommand
		isCreate bool
	}{
		"valid create": {
			command:  validReceiptCommand(),
			isCreate: true,
		},
		"valid update": {
			command:  validReceiptCommand(),
			isCreate: false,
		},
		"valid with nested entities": {
			command: UpsertReceiptCommand{
				Name:         "Test Receipt",
				Amount:       decimal.NewFromFloat(25.50),
				Date:         time.Now(),
				GroupId:      1,
				PaidByUserID: 1,
				Status:       models.ReceiptStatus("OPEN"),
				Categories:   []UpsertCategoryCommand{{Name: "Food"}},
				Tags:         []UpsertTagCommand{{Name: "Groceries"}},
				Items: []UpsertItemCommand{{
					Amount: decimal.NewFromFloat(10.00),
					Name:   "Item 1",
					Status: models.ITEM_OPEN,
				}},
				Comments: []UpsertCommentCommand{{
					Comment: "A comment",
					UserId:  &userId,
				}},
			},
			isCreate: true,
		},
		"zero amount": {
			command: UpsertReceiptCommand{
				Name:         "Test Receipt",
				Amount:       decimal.Zero,
				Date:         time.Now(),
				GroupId:      1,
				PaidByUserID: 1,
				Status:       models.ReceiptStatus("OPEN"),
			},
			isCreate: true,
		},
		"negative amount for refund": {
			command: UpsertReceiptCommand{
				Name:         "Store Return",
				Amount:       decimal.NewFromFloat(-25.50),
				Date:         time.Now(),
				GroupId:      1,
				PaidByUserID: 1,
				Status:       models.ReceiptStatus("OPEN"),
			},
			isCreate: true,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			vErr := test.command.Validate(1, test.isCreate)

			if len(vErr.Errors) > 0 {
				utils.PrintTestError(t, len(vErr.Errors), 0)
			}
		})
	}
}

func TestUpsertReceiptCommand_Validate_MissingFields(t *testing.T) {
	tests := map[string]struct {
		modify        func(cmd *UpsertReceiptCommand)
		expectedError string
	}{
		"missing name": {
			modify:        func(cmd *UpsertReceiptCommand) { cmd.Name = "" },
			expectedError: "name",
		},
		"missing date": {
			modify:        func(cmd *UpsertReceiptCommand) { cmd.Date = time.Time{} },
			expectedError: "date",
		},
		"missing groupId": {
			modify:        func(cmd *UpsertReceiptCommand) { cmd.GroupId = 0 },
			expectedError: "groupId",
		},
		"missing paidByUserId": {
			modify:        func(cmd *UpsertReceiptCommand) { cmd.PaidByUserID = 0 },
			expectedError: "paidByUserId",
		},
		"missing status": {
			modify:        func(cmd *UpsertReceiptCommand) { cmd.Status = "" },
			expectedError: "status",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			cmd := validReceiptCommand()
			test.modify(&cmd)

			vErr := cmd.Validate(1, true)

			if _, exists := vErr.Errors[test.expectedError]; !exists {
				utils.PrintTestError(t, "error should exist for field", test.expectedError)
			}
		})
	}
}

func TestUpsertReceiptCommand_Validate_NestedCategoryErrors(t *testing.T) {
	cmd := validReceiptCommand()
	cmd.Categories = []UpsertCategoryCommand{{Name: ""}}

	vErr := cmd.Validate(1, true)

	if _, exists := vErr.Errors["categories.0.name"]; !exists {
		utils.PrintTestError(t, "error should exist for field", "categories.0.name")
	}
}

func TestUpsertReceiptCommand_Validate_NestedTagErrors(t *testing.T) {
	cmd := validReceiptCommand()
	cmd.Tags = []UpsertTagCommand{{Name: ""}}

	vErr := cmd.Validate(1, true)

	if _, exists := vErr.Errors["tags.0.name"]; !exists {
		utils.PrintTestError(t, "error should exist for field", "tags.0.name")
	}
}

func TestUpsertReceiptCommand_Validate_NestedItemErrors(t *testing.T) {
	cmd := validReceiptCommand()
	cmd.Items = []UpsertItemCommand{{}}

	vErr := cmd.Validate(1, true)

	if _, exists := vErr.Errors["receiptItems.0.name"]; !exists {
		utils.PrintTestError(t, "error should exist for field", "receiptItems.0.name")
	}
}

func TestUpsertReceiptCommand_Validate_ItemsAgainstDocumentAmount(t *testing.T) {
	documentAmount := decimal.NewFromInt(100)
	cmd := validReceiptCommand()
	cmd.Amount = decimal.NewFromInt(154)
	cmd.DocumentAmount = &documentAmount
	cmd.DocumentCurrencyCode = "USD"
	cmd.Items = []UpsertItemCommand{{
		Name: "Printed line item", Amount: decimal.NewFromInt(120), Status: models.ITEM_OPEN,
	}}

	if _, exists := cmd.Validate(1, true).Errors["receiptItems.0.amount"]; !exists {
		t.Fatal("expected item amount to be compared with documentAmount")
	}
}

func TestUpsertReceiptCommand_Validate_InvalidDocumentCurrency(t *testing.T) {
	cmd := validReceiptCommand()
	cmd.DocumentCurrencyCode = "NOT-A-CURRENCY"
	if _, exists := cmd.Validate(1, true).Errors["documentCurrencyCode"]; !exists {
		t.Fatal("expected invalid ISO 4217 code error")
	}
}

func TestUpsertReceiptCommand_Validate_NestedCommentErrors(t *testing.T) {
	cmd := validReceiptCommand()
	cmd.Comments = []UpsertCommentCommand{{}}

	vErr := cmd.Validate(1, true)

	if _, exists := vErr.Errors["comments.0.comment"]; !exists {
		utils.PrintTestError(t, "error should exist for field", "comments.0.comment")
	}
}

func TestUpsertReceiptCommand_Validate_EmptyCommand(t *testing.T) {
	command := UpsertReceiptCommand{}

	vErr := command.Validate(1, true)

	// name, date, groupId, paidByUserId, status
	if len(vErr.Errors) < 5 {
		utils.PrintTestError(t, len(vErr.Errors), "at least 5")
	}
}
