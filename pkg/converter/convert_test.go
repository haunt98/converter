package converter

import "testing"

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
			want: "a__b_c__d",
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