package services

import "testing"

func TestNormaliseSupplierName(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "trim and fold", input: "  GitHub  ", expected: "github"},
		{name: "collapse whitespace", input: "GitHub    Inc", expected: "github inc"},
		{name: "strip punctuation", input: "GitHub, Inc.", expected: "github inc"},
		{name: "case variant", input: "GITHUB INC", expected: "github inc"},
		{name: "quotes and parens", input: `"GitHub" (Inc)`, expected: "github inc"},
		{name: "empty", input: "   ", expected: ""},
		{name: "punctuation only", input: "...", expected: ""},
		{name: "domain keeps letters", input: "github.com", expected: "githubcom"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NormaliseSupplierName(tc.input)
			if actual != tc.expected {
				t.Fatalf("NormaliseSupplierName(%q) = %q, want %q", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestUniqueNormalisedAliasesDedupes(t *testing.T) {
	t.Parallel()

	aliases := uniqueNormalisedAliases([]string{"GitHub, Inc.", "github inc", "  ", "github.com"})
	if len(aliases) != 2 {
		t.Fatalf("expected 2 unique aliases, got %#v", aliases)
	}
	if aliases[0].NormalisedName != "github inc" {
		t.Fatalf("first alias = %q", aliases[0].NormalisedName)
	}
	if aliases[1].NormalisedName != "githubcom" {
		t.Fatalf("second alias = %q", aliases[1].NormalisedName)
	}
}
