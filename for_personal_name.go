package kanji

import (
	"unicode"
)

// IsForPersonalNames returns true if the rune r is kanji designated for personal names.
func IsForPersonalNames(r rune) bool {
	return IsRegularUse(r) || unicode.Is(DesignatedForPersonalNamesRangeTable, r)
}

// IsNotForPersonalNames returns true if the rune r is in the `unicode.Han` and is not kanji designated for personal names.
func IsNotForPersonalNames(r rune) bool {
	return unicode.Is(unicode.Han, r) && !IsForPersonalNames(r)
}
