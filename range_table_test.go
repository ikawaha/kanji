package kanji

import (
	"fmt"
	"testing"
	"unicode"
)

type range16 unicode.Range16

func (r range16) String() string {
	return fmt.Sprintf("{Lo:0x%04X, Hi:0x%04X, Stride:%d}", r.Lo, r.Hi, r.Stride)
}

type range32 unicode.Range32

func (r range32) String() string {
	return fmt.Sprintf("{Lo:0x%0X, Hi:0x%0X, Stride:%d}", r.Lo, r.Hi, r.Stride)
}

func Test_CheckingRangeTables(t *testing.T) {
	// R16
	for _, table := range []*unicode.RangeTable{
		StandardRangeTable,
		OldFormRangeTable,
		TolerableRangeTable,
		DesignatedForPersonalNamesRangeTable,
	} {
		for i, v := range table.R16 {
			// stride check
			if want, got := uint16(1), v.Stride; want != got {
				t.Errorf("stride want %d, got %d", want, got)
			}
			// range check
			if v.Lo > v.Hi {
				t.Errorf("range error, %v", range16(v))
			}
			// overlap check
			if i != 0 {
				vv := table.R16[i-1]
				if vv.Hi > v.Lo {
					t.Errorf("overlap %v, %v", range16(vv), range16(v))
				}
			}
			// boundary check
			if !unicode.Is(table, rune(v.Lo)) {
				t.Errorf("boundary value error (lo), %v, 0x%04X", range16(v), v.Lo)
			}
			if !unicode.Is(table, rune(v.Hi)) {
				t.Errorf("boundary value error (hi), %v, 0x%04X", range16(v), v.Hi)
			}
		}
		// R32
		for i, v := range table.R32 {
			// stride check
			if want, got := uint32(1), v.Stride; want != got {
				t.Errorf("stride want %d, got %d", want, got)
			}
			// overlap check
			if i != 0 {
				vv := table.R32[i-1]
				if vv.Hi > v.Lo {
					t.Errorf("overlap %v, %v", range32(vv), range32(v))
				}
			}
			// boundary check
			if !unicode.Is(table, rune(v.Lo)) {
				t.Errorf("boundary value error, %v, 0x%0X", range32(v), v.Lo)
			}
			if !unicode.Is(table, rune(v.Hi)) {
				t.Errorf("boundary value error, %v, 0x%0X", range32(v), v.Hi)
			}
		}
	}
}
