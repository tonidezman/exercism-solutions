package bob

import (
	"regexp"
	"strings"
	"unicode"
)

func trimWhiteSpace(response string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, response)
}

// Hey returns remark
func Hey(remark string) string {
	r := trimWhiteSpace(remark)
	questionChar := "?"
	emptyString := ""

	if r == emptyString {
		return "Fine. Be that way!"
	} else if isLastChar(r, questionChar) && atLeastOneWordChar(r) && isUpper(r) {
		return "Calm down, I know what I'm doing!"
	} else if isUpper(r) && atLeastOneWordChar(r) {
		return "Whoa, chill out!"
	} else if isLastChar(r, questionChar) {
		return "Sure."
	}
	return "Whatever."
}

func atLeastOneWordChar(remark string) bool {
	result, _ := regexp.MatchString("[a-zA-z]+", remark)
	return result
}

func isUpper(remark string) bool {
	return strings.ToUpper(remark) == remark
}

func isLastChar(remark, char string) bool {
	lastChar := string(remark[len(remark)-1])
	if lastChar == char {
		return true
	}
	return false
}
