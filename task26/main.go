package main

import (
	"fmt"
	"strings"
)

func main() {
	t1 := "abcd" // true
	t2 := "abCdefAaf" // false
	t3 := "aabcd" // false

	// Check and print results
	fmt.Println(IsUnique(t1))
	fmt.Println(IsUnique(t2))
	fmt.Println(IsUnique(t3))
}

func IsUnique(s string) bool {
	s = strings.ToLower(s)

	charMap := make(map[rune]bool)

	for _, char := range s {
		if _, exists := charMap[char]; exists {
			return false
		}
		charMap[char] = true
	}

	return true
}
