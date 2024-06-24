package modifiers

import (
	"regexp"
	"strings"
)

func HandleQuotes(text string) string {
	reSingleQuotes := regexp.MustCompile(`'\s*(.*?)\s*'`)
	text = reSingleQuotes.ReplaceAllStringFunc(text, func(match string) string {
		matches := reSingleQuotes.FindStringSubmatch(match)
		return "'" + strings.TrimSpace(matches[1]) + "'"
	})

	// Remove trailing space if it exists
	text = strings.TrimSpace(text)

	return text
}
