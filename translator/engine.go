package translator

import (
	"strings"
)

func Translate(input string) string {
	for {
		if !beginsWithVowel(input) {
			input = moveLetterToEnd(input)
		} else {
			break
		}
	}
	return input + "ay"
}

func moveLetterToEnd(input string) string {
	firstLetter := string(input[0])
	input += firstLetter
	return string(input[1:])
}

func beginsWithVowel(input string) bool {
	if strings.HasPrefix(strings.ToUpper(input), "A") ||
		strings.HasPrefix(strings.ToUpper(input), "E") ||
		strings.HasPrefix(strings.ToUpper(input), "I") ||
		strings.HasPrefix(strings.ToUpper(input), "O") ||
		strings.HasPrefix(strings.ToUpper(input), "U") {
		return true
	} else {
		return false
	}
}
