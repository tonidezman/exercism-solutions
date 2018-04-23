package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate returns Acronym for the string
func Abbreviate(s string) string {
	var acronym string
	words := regexp.MustCompile("[ -]").Split(s, -1)
	for _, word := range words {
		acronym += string(word[0])
	}
	return strings.ToUpper(acronym)
}
