package goleetcode

import (
	"unicode"
)

func findWords(words []string) []string {
	kb := map[rune]int{
		'Q': 1, 'W': 1, 'E': 1, 'R': 1, 'T': 1, 'Y': 1, 'U': 1, 'I': 1, 'O': 1, 'P': 1,
		'A': 2, 'S': 2, 'D': 2, 'F': 2, 'G': 2, 'H': 2, 'J': 2, 'K': 2, 'L': 2,
		'Z': 3, 'X': 3, 'C': 3, 'V': 3, 'B': 3, 'N': 3, 'M': 3,
	}
	result := []string{}
	for _, word := range words {
		if word == "" {
			continue
		}

		row := kb[unicode.ToUpper(rune(word[0]))]
		for i := 1; i < len(word); i++ {
			if kb[unicode.ToUpper(rune(word[i]))] != row {
				goto outter
			}
		}
		result = append(result, word)
	outter:
	}
	return result
}
