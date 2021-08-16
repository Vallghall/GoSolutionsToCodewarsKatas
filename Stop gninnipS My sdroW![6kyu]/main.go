package main

import (
	"fmt"
	"strings"
)

func SpinWords(str string) string {
	words := strings.Fields(str)
	sb := strings.Builder{}

	for _, word := range words {
		if len(word) >= 5 {
			sb.WriteString(wordReverse(word, len(word)) + " ")
		} else {
			sb.WriteString(word + " ")
		}
	}
	return strings.TrimRight(sb.String(), " ")
}

func wordReverse(word string, length int) string {
	wordReversed := make([]rune, length, length)
	for i, letter := range word {
		wordReversed[length-1-i] = letter
	}
	return string(wordReversed)
}

func main() {
	fmt.Println(SpinWords("Checking work check"))
}
