package commands

import "testing"

func TestUpsertSupplierProfileCommandValidate_RequiresNameAndDefault(t *testing.T) {
	command := UpsertSupplierProfileCommand{}
	vErr := command.Validate()
	if vErr.Errors["name"] == "" {
		t.Fatal("expected name required")
	}
	if vErr.Errors["defaults"] == "" {
		t.Fatal("expected at least one default")
	}
}

func TestUpsertSupplierProfileCommandValidate_AcceptsCurrencyDefault(t *testing.T) {
	code := "usd"
	command := UpsertSupplierProfileCommand{Name: "GitHub", ExpectedDocumentCurrencyCode: &code}
	vErr := command.Validate()
	if len(vErr.Errors) != 0 {
		t.Fatalf("unexpected errors: %#v", vErr.Errors)
	}
	if command.ExpectedDocumentCurrencyCode == nil || *command.ExpectedDocumentCurrencyCode != "USD" {
		t.Fatalf("expected currency normalised to USD, got %#v", command.ExpectedDocumentCurrencyCode)
	}
}

func TestUpsertSupplierProfileCommandValidate_RejectsInvalidCurrency(t *testing.T) {
	code := "XXXX"
	command := UpsertSupplierProfileCommand{Name: "GitHub", ExpectedDocumentCurrencyCode: &code}
	vErr := command.Validate()
	if vErr.Errors["expectedDocumentCurrencyCode"] == "" {
		t.Fatal("expected invalid currency error")
	}
}

func TestResolveSupplierProfileCommandValidate_RequiresName(t *testing.T) {
	command := ResolveSupplierProfileCommand{}
	vErr := command.Validate()
	if vErr.Errors["name"] == "" {
		t.Fatal("expected name required")
	}
}
