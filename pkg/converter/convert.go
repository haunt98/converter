package converter

import "strings"

func Convert(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, "'", "_")
	text = strings.ReplaceAll(text, `"`, "_")
	text = strings.ReplaceAll(text, " ", "_")
	text = strings.ReplaceAll(text, ",", "_")
	text = strings.ReplaceAll(text, "(", "_")
	text = strings.ReplaceAll(text, ")", "_")

	return text
}
