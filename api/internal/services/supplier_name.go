package services

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/unicode/norm"
)

// presentationPunctuation is a conservative set of decorative marks stripped
// during supplier-name matching. Hyphens and ampersands are kept so they can
// still distinguish names.
var presentationPunctuation = map[rune]struct{}{
	',': {}, '.': {}, ';': {}, ':': {}, '!': {}, '?': {},
	'(': {}, ')': {}, '[': {}, ']': {}, '{': {}, '}': {},
	'"': {}, '\'': {}, '`': {}, '“': {}, '”': {}, '‘': {}, '’': {},
	'·': {}, '•': {}, '™': {}, '®': {}, '©': {},
}

// NormaliseSupplierName trims, Unicode-folds, strips presentation punctuation,
// and collapses whitespace. The display name is stored separately; this value
// is what matching and collision checks use.
func NormaliseSupplierName(name string) string {
	s := strings.TrimSpace(name)
	if len(s) == 0 {
		return ""
	}

	s = norm.NFKC.String(s)
	s = cases.Fold().String(s)

	var builder strings.Builder
	builder.Grow(len(s))
	prevSpace := false
	for _, r := range s {
		if _, ok := presentationPunctuation[r]; ok {
			continue
		}
		if unicode.IsSpace(r) {
			if !prevSpace && builder.Len() > 0 {
				builder.WriteRune(' ')
				prevSpace = true
			}
			continue
		}
		builder.WriteRune(r)
		prevSpace = false
	}

	return strings.TrimSpace(builder.String())
}

type normalisedAlias struct {
	Name           string
	NormalisedName string
}

func uniqueNormalisedAliases(aliases []string) []normalisedAlias {
	seen := make(map[string]struct{})
	result := make([]normalisedAlias, 0, len(aliases))

	for _, alias := range aliases {
		trimmed := strings.TrimSpace(alias)
		normalised := NormaliseSupplierName(trimmed)
		if len(normalised) == 0 {
			continue
		}
		if _, exists := seen[normalised]; exists {
			continue
		}
		seen[normalised] = struct{}{}
		result = append(result, normalisedAlias{Name: trimmed, NormalisedName: normalised})
	}

	return result
}
