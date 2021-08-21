package decode_morse_code

import (
	mt "github.com/Vallghall/GoSolutionsToCodewarsKatas/decode_morse_code/morse_table"
	"strings"
)

func DecodeMorse(morseCode string) string {
	sb := strings.Builder{}
	for _, word := range strings.Split(morseCode, "   ") {
		for _, letter := range strings.Fields(word) {
			sb.WriteString(mt.MORSE_CODE[letter])
		}
		sb.WriteString(" ")
	}
	return strings.TrimSpace(sb.String())
}
