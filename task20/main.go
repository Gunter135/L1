package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ReverseWords("fish or not to fish"))
}
// TODO:
// reverse words with specified delimiter

func ReverseWords(str string) string {
	// If we are looking to split strings with whitespaces only
	words := strings.Fields(str)
	var result strings.Builder
	for i := len(words) - 1; i > 0; i-- {
		result.WriteString(words[i])
		result.WriteString(" ")
	}
	result.WriteString(words[0])
	return result.String()
}
