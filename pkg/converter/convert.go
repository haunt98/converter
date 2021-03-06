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
	tempText := ""
	for _, c := range text {
		if isNormalChar(c) {
			tempText += string(c)
			continue
		}

		tempText += string(normalizedChar)
	}
	text = tempText

	// remove duplicate _
	// a__b -> a_b
	text = reNormalizedDuplicate.ReplaceAllString(text, string(normalizedChar))

	// remove prefix, postfix _
	// _abc_ -> abc
	text = strings.TrimPrefix(text, string(normalizedChar))
	text = strings.TrimSuffix(text, string(normalizedChar))

	return text
}

func isNormalText(text string) bool {
	for _, c := range text {
		if !isNormalChar(c) {
			return false
		}
	}

	return true
}

func isNormalChar(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= '0' && c <= '9') ||
		c == normalizedChar
}
