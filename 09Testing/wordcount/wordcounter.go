package wordcount

import "strings"

func CountWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}
