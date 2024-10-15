package main

import "fmt"

func removeElement(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

func main() {
	slice := []int{10, 20, 30, 40, 50}

	indexToRemove := 2
	newSlice := removeElement(slice, indexToRemove)

	fmt.Println("Old Slice:", slice)
	fmt.Println("New Slice:", newSlice)
}
