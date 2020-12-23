package converter

import (
	"regexp"
	"strings"
)

const (
	normalizedChar = '_'
)

var reNormalizedDuplicate = regexp.MustCompile(`_+`)

func Convert(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	// whitelist approach
	finalText := ""
	for _, c := range text {
		if isNormalRune(c) {
			finalText += string(c)
			continue
		}

		finalText += string(normalizedChar)
	}
	text = finalText

	// remove duplicate _
	text = reNormalizedDuplicate.ReplaceAllString(text, string(normalizedChar))

	return text
}

func isNormalText(text string) bool {
	for _, c := range text {
		if !isNormalRune(c) {
			return false
		}
	}

	return true
}

func isNormalRune(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= '0' && c <= '9') ||
		c == normalizedChar
}
