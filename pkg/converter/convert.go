package converter

import (
	"regexp"
	"strings"
)

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

var reNormalizedDuplicate = regexp.MustCompile(`_+`)

func Convert(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	for _, denyChar := range denyChars {
		text = strings.ReplaceAll(text, denyChar, normalizedChar)
	}

	text = reNormalizedDuplicate.ReplaceAllString(text, "_")

	return text
}
