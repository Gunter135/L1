package main

import "fmt"

func main() {
	x := 5 //101
	y := 6 //110
	fmt.Printf("Before swap: x=%d,y=%d\n", x, y)
	x = x ^ y // 101 XOR 110 = 011
	y = x ^ y // 011 XOR 110 = 101
	x = x ^ y // 011 XOR 101 = 110
	fmt.Printf("After swap x=%d,y=%d\n", x, y)
	// Another way to swap two numbers
	m := 5
	n := 6
	fmt.Printf("Before swap: m=%d,n=%d\n", m, n)
	m += n // m = 11 n = 6
	n = m - n // m = 11 n = 5
	m = m - n // m = 6 n = 5
	fmt.Printf("After swap m=%d,n=%d\n", m, n)
}