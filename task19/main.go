package main

import "fmt"

func main() {
	fmt.Println(ReverseString("Гру"))
}

func ReverseString(str string) string {
	rs := []rune(str)
	size := len(rs)
	for i := 0; i < size/2; i++ {
		rs[i], rs[size-1-i] = rs[size-1-i], rs[i]
		fmt.Println(string(rs))
	}
	return string(rs)
}
