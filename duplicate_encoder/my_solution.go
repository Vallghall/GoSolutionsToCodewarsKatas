package duplicate_encoder

import (
	"strings"
	"unicode"
)

func DuplicateEncode(word string) string {
	letterCounter := make(map[rune]int)
	sb := strings.Builder{}
	for _, letter := range word {
		letterCounter[unicode.ToLower(letter)]++
	}
	for _, letter := range word {
		if letterCounter[unicode.ToLower(letter)] > 1 {
			sb.WriteString(")")
		} else {
			sb.WriteString("(")
		}
	}
	return sb.String()
}
