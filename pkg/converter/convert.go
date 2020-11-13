package converter

import "strings"

const (
	normalizedChar = "_"
)

var denyChars = []string{
	`"`,
	"'",
	" ",
	"\t",
	"\n",
	",",
	"(",
	")",
	":",
}

func Convert(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	for _, denyChar := range denyChars {
		text = strings.ReplaceAll(text, denyChar, normalizedChar)
	}

	return text
}
