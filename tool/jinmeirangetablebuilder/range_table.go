package main

import (
	"fmt"
	"io"
	"sort"
	"unicode"
)

// MaxUint16 is the maximum number of the type uint16 (1<<16-1)
const MaxUint16 = 1<<16 - 1

// RangeTable returns a range table for specified runes.
func RangeTable(runes []rune) *unicode.RangeTable {
	if len(runes) == 0 {
		return nil
	}
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	var ret unicode.RangeTable
	for start := 0; start < len(runes); {
		end := start
		for j := start + 1; j < len(runes); j++ {
			if runes[j-1]+1 == runes[j] {
				end = j
				continue
			}
			break
		}
		if runes[end] <= MaxUint16 {
			ret.R16 = append(ret.R16, unicode.Range16{
				Lo:     uint16(runes[start]),
				Hi:     uint16(runes[end]),
				Stride: 1,
			})
		} else {
			ret.R32 = append(ret.R32, unicode.Range32{
				Lo:     uint32(runes[start]),
				Hi:     uint32(runes[end]),
				Stride: 1,
			})
		}
		start = end + 1
	}
	return &ret
}

// DumpRangeTable write out the range table in Go source code format.
func DumpRangeTable(w io.Writer, name string, table *unicode.RangeTable) {
	fmt.Fprintf(w, "var %s = &unicode.RangeTable{\n", name)
	if len(table.R16) > 0 {
		fmt.Println("\tR16: []unicode.Range16{")
		for _, v := range table.R16 {
			fmt.Fprintf(w, "\t\t{Lo: 0x%X, Hi: 0x%X, Stride: 1},\n", v.Lo, v.Hi)
		}
		fmt.Fprintln(w, "\t},")
	}
	if len(table.R32) > 0 {
		fmt.Println("\tR32: []unicode.Range32{")
		for _, v := range table.R32 {
			fmt.Fprintf(w, "\t\t{Lo: 0x%X, Hi: 0x%X, Stride: 1},\n", v.Lo, v.Hi)
		}
		fmt.Fprintln(w, "\t},")
	}
	fmt.Fprintln(w, "}")
}
