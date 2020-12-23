package converter

import (
	"regexp"
	"strings"
)

const (
	normalizedChar = "_"
)

var reNormalizedDuplicate = regexp.MustCompile(`_+`)

func Convert(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	// whitelist approach
	finalText := ""
	for _, c := range text {
		if (c >= 'a' && c <= 'z') ||
			(c >= '0' && c <= '9') {
			finalText += string(c)
			continue
		}

		finalText += "_"
	}
	text = finalText

	// remove duplicate _
	text = reNormalizedDuplicate.ReplaceAllString(text, "_")

	return text
}
