package kanji

import (
	"unicode"
)

// IsHan returns true if the rune is in `unicode.Han`.
func IsHan(r rune) bool {
	return unicode.Is(unicode.Han, r)
}
