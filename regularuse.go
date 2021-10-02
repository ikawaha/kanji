package kanji

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// IsRegularUse returns true if the rune r is in the regular-use list (常用漢字表).
func IsRegularUse(r rune) bool {
	return unicode.In(r, StandardRangeTable, OldFormRangeTable, TolerableRangeTable)
}

// IsStandardRegularUse returns true if the rune r is in the regular-use standard list (常用漢字表：標準字体).
func IsStandardRegularUse(r rune) bool {
	return unicode.Is(StandardRangeTable, r)
}

// IsOldFormRegularUse returns true if the rune r is in the regular-use old form list (常用漢字表：旧字体).
func IsOldFormRegularUse(r rune) bool {
	return unicode.Is(OldFormRangeTable, r)
}

// IsTolerableRegularUse returns true if the rune r is in the regular-use old form list (常用漢字表：許容字体).
func IsTolerableRegularUse(r rune) bool {
	return unicode.Is(TolerableRangeTable, r)
}

// IsNotRegularUse returns true if the rune r is in the `unicode.Han` and is not in the regular use list (常用漢字表).
func IsNotRegularUse(r rune) bool {
	return unicode.Is(unicode.Han, r) && !IsRegularUse(r)
}

// ReplaceNotRegularUseAll returns a copy of the string s with each run of not in regular-use kanji
// replaced by the replacement string, which may be empty.
func ReplaceNotRegularUseAll(s, replacement string) string {
	return replace(s, replacement, func(r rune) bool { return !IsNotRegularUse(r) })
}

// Option represents an option for the discriminator.
type Option func(d *RegularUseDiscriminator)

// Allow is a discriminator option to set allow characters.
func Allow(r ...rune) Option {
	return func(d *RegularUseDiscriminator) {
		d.allow = append(d.allow, r...)
	}
}

// Disallow is a discriminator option to set disallow characters.
func Disallow(r ...rune) Option {
	return func(d *RegularUseDiscriminator) {
		d.disallow = append(d.disallow, r...)
	}
}

// RegularUseDiscriminator determines if a character is in regular-use kanji or allowed/disallowed character.
type RegularUseDiscriminator struct {
	allow    []rune
	disallow []rune
}

// NewRegularUseDiscriminator returns a regular kanji character discriminator.
func NewRegularUseDiscriminator(options ...Option) *RegularUseDiscriminator {
	var ret RegularUseDiscriminator
	for _, option := range options {
		option(&ret)
	}
	return &ret
}

// IsNotRegularUse returns true if the rune r is in disallowed characters, else if return false r is in allowed characters,
// otherwise whether r is not in regular-use kanji.
func (d *RegularUseDiscriminator) IsNotRegularUse(r rune) bool {
	for _, v := range d.allow {
		if v == r {
			return false
		}
	}
	for _, v := range d.disallow {
		if v == r {
			return true
		}
	}
	return IsNotRegularUse(r)
}

// ReplaceNotRegularUseAll returns a copy of the string s with each run of not in regular-use kanji
// replaced by the replacement string, which may be empty.
func (d *RegularUseDiscriminator) ReplaceNotRegularUseAll(s, replacement string) string {
	return replace(s, replacement, func(r rune) bool { return !d.IsNotRegularUse(r) })
}

func replace(s, replacement string, is func(rune) bool) string {
	var b strings.Builder

	for i, c := range s {
		if c != utf8.RuneError && is(c) {
			continue
		}
		if !is(c) {
			b.Grow(len(s) + len(replacement))
			b.WriteString(s[:i])
			s = s[i:]
			break
		}
		_, wid := utf8.DecodeRuneInString(s[i:])
		if wid == 1 {
			b.Grow(len(s) + len(replacement))
			b.WriteString(s[:i])
			s = s[i:]
			break
		}
	}

	// Fast path for unchanged input
	if b.Cap() == 0 { // didn't call b.Grow above
		return s
	}

	invalid := false // previous byte was from an invalid UTF-8 sequence
	for i := 0; i < len(s); {
		c := s[i]
		if c < utf8.RuneSelf {
			i++
			invalid = false
			b.WriteByte(c)
			continue
		}
		r, wid := utf8.DecodeRuneInString(s[i:])
		if wid == 1 {
			i++
			if !invalid {
				invalid = true
				b.WriteString(replacement)
			}
			continue
		}
		invalid = false
		if is(r) {
			b.WriteString(s[i : i+wid])
		} else {
			b.WriteString(replacement)
		}
		i += wid
	}
	return b.String()
}
