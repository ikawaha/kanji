package kanji

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

func TestIsForPersonalNames(t *testing.T) {
	f, err := os.Open("./testdata/golden_jinmei.txt")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Scan()
	line := s.Text()
	if !strings.HasPrefix(line, "!!!") {
		t.Fatalf("invalid file format, %s", line)
	}
	t.Log(line)
	i := 1
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, "!!!") {
			t.Log(line)
			break
		}
		runes := []rune(line)
		if len(runes) < 1 {
			log.Println("empty line, line no:", i)
			continue
		}
		for _, v := range runes {
			if v == '‐' {
				continue
			}
			if got, want := IsForPersonalNames(v), true; got != want {
				t.Errorf("IsForPersonalNames(%c)=%v, want %v", v, got, want)
			}
		}
	}
	if err := s.Err(); err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	for s.Scan() {
		line := s.Text()
		runes := []rune(line)
		if len(runes) < 1 {
			log.Println("empty line, line no:", i)
			continue
		}
		if got, want := IsForPersonalNames(runes[0]), true; got != want {
			t.Errorf("IsForPersonalNames(%c)=%v, want %v", runes[0], got, want)
		}
	}
}

func TestIsNotForPersonalNames(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "OK",
			args: "漢字以外のひらがなやカタカナや😀などもOKとしています!",
			want: false,
		},
		{
			name: "for personal names",
			args: "紗也茄",
			want: false,
		},
		{
			name: "NG",
			args: "棗薔薇玻繚茗厦祟",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.args {
				if got := IsNotForPersonalNames(v); got != tt.want {
					t.Errorf("IsNotForPersonalNames(%c) = %v, want %v", v, got, tt.want)
				}
			}
		})
	}
}
