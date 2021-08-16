package main

import (
	"fmt"
)

func ValidParentheses(parens string) bool {
	if len(parens)%2 != 0 {
		return false
	}

	parensCounter := 0
	for _, val := range parens {
		switch val {
		case '(':
			parensCounter++
		case ')':
			parensCounter--
		default:
			continue
		}
		if parensCounter == -1 {
			return false
		}
	}
	if parensCounter != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(ValidParentheses("()()()())"))
}
