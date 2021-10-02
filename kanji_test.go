package kanji

import (
	"testing"
)

func TestIsHan(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Hiragana",
			args: args{
				r: 'あ',
			},
			want: false,
		},
		{
			name: "Katakana",
			args: args{
				r: 'カ',
			},
			want: false,
		},
		{
			name: "Kanji",
			args: args{
				r: '漢',
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHan(tt.args.r); got != tt.want {
				t.Errorf("IsHan() = %v, want %v", got, tt.want)
			}
		})
	}
}
