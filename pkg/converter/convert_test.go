package converter

import (
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

const (
	maxFuzzyTest = 1000
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "no change",
			text: "abc",
			want: "abc",
		},
		{
			name: "single quote",
			text: "'abc'",
			want: "_abc_",
		},
		{
			name: "double quote",
			text: `"abc"`,
			want: "_abc_",
		},
		{
			name: "space",
			text: "ab cd",
			want: "ab_cd",
		},
		{
			name: "mixed",
			text: `a"'b c"'d`,
			want: "a_b_c_d",
		},
		{
			name: "comma",
			text: "a,b",
			want: "a_b",
		},
		{
			name: "left parenthesis",
			text: "a(b",
			want: "a_b",
		},
		{
			name: "right parenthesis",
			text: "a)b",
			want: "a_b",
		},
		{
			name: "parenthesis",
			text: "(ab)",
			want: "_ab_",
		},
		{
			name: "tab",
			text: "a\tb\tc",
			want: "a_b_c",
		},
		{
			name: "newline",
			text: "a\nb\nc",
			want: "a_b_c",
		},
		{
			name: "colon",
			text: "a:b:c",
			want: "a_b_c",
		},
		{
			name: "_ duplicate",
			text: "a__b___c",
			want: "a_b_c",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Convert(tc.text)
			if got != tc.want {
				t.Errorf("got %s want %s", got, tc.want)
			}
		})
	}
}

func TestFuzzyConvert(t *testing.T) {
	f := fuzz.New()
	var text string

	for i := 0; i < maxFuzzyTest; i++ {
		f.Fuzz(&text)
		text = Convert(text)
		assert.Equal(t, true, isNormalText(text))
	}
}
