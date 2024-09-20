package main

import "fmt"

func main() {
	fmt.Println(MakeSet([]string{"cat", "cat", "dog", "cat", "tree"}))
}

func MakeSet(strings []string) map[string]struct{} {
	result := make(map[string]struct{})
	for _, v := range strings {
		result[v] = struct{}{}
	}
	return result
}
