package commands

import (
	"encoding/json"
	"net/http"
	"strings"

	"golang.org/x/text/currency"
	"receipt-wrangler/api/internal/models"
	"receipt-wrangler/api/internal/structs"
	"receipt-wrangler/api/internal/utils"
)

type UpsertSupplierProfileCommand struct {
	Name                         string   `json:"name"`
	Aliases                      []string `json:"aliases"`
	CategoryIds                  []uint   `json:"categoryIds"`
	TagIds                       []uint   `json:"tagIds"`
	ExpectedDocumentCurrencyCode *string  `json:"expectedDocumentCurrencyCode"`
	Enabled                      *bool    `json:"enabled"`
}

func (command *UpsertSupplierProfileCommand) LoadDataFromRequest(w http.ResponseWriter, r *http.Request) error {
	bytes, err := utils.GetBodyData(w, r)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, &command)
}

func (command *UpsertSupplierProfileCommand) Validate() structs.ValidatorError {
	errors := make(map[string]string)
	vErr := structs.ValidatorError{}

	if len(strings.TrimSpace(command.Name)) == 0 {
		errors["name"] = "Name is required"
	}

	hasCurrency := false
	if command.ExpectedDocumentCurrencyCode != nil {
		code := strings.ToUpper(strings.TrimSpace(*command.ExpectedDocumentCurrencyCode))
		if len(code) == 0 {
			command.ExpectedDocumentCurrencyCode = nil
		} else {
			command.ExpectedDocumentCurrencyCode = &code
			if _, err := currency.ParseISO(code); err != nil {
				errors["expectedDocumentCurrencyCode"] = "Expected Document Currency Code must be a valid ISO 4217 code"
			} else {
				hasCurrency = true
			}
		}
	}

	if len(command.CategoryIds) == 0 && len(command.TagIds) == 0 && !hasCurrency {
		errors["defaults"] = "At least one default is required (categories, tags, or expected currency)"
	}

	vErr.Errors = errors
	return vErr
}

type ResolveSupplierProfileCommand struct {
	Name string `json:"name"`
}

func (command *ResolveSupplierProfileCommand) LoadDataFromRequest(w http.ResponseWriter, r *http.Request) error {
	bytes, err := utils.GetBodyData(w, r)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, &command)
}

func (command *ResolveSupplierProfileCommand) Validate() structs.ValidatorError {
	errors := make(map[string]string)
	vErr := structs.ValidatorError{}

	if len(strings.TrimSpace(command.Name)) == 0 {
		errors["name"] = "Name is required"
	}

	vErr.Errors = errors
	return vErr
}

type ResolveSupplierProfileResponse struct {
	Profile *models.SupplierProfile `json:"profile"`
}
