package kanji

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestNewOldFormNewFormReplacer(t *testing.T) {
	f, err := os.Open("./testdata/golden_old-new.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer f.Close()

	replacer := NewOldFormNewFormReplacer()

	s := bufio.NewScanner(f)
	var line int
	for s.Scan() {
		l := s.Text()
		line++
		if strings.HasPrefix(l, "!") {
			continue
		}
		a := strings.Split(l, " ")
		if len(a) != 4 {
			t.Errorf("invalid golden file, line=%d, %s", line, l)
			continue
		}
		code, err := strconv.ParseInt(a[0], 16, strconv.IntSize)
		if err != nil {
			t.Errorf("invalid golden file, line=%d, parse int error %v, %s", line, err, l)
			continue
		}
		r := rune(code)
		t.Run(l, func(t *testing.T) {
			got := replacer.Replace(string(r))
			want := a[3]
			if got != want {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
	if err := s.Err(); err != nil {
		t.Errorf("unexpected error, %v", err)
	}
}
