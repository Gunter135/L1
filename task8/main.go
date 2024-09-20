package main

import "fmt"

func main() {
	var x int64 = 10 // 1010
	fmt.Println(x)
	x |= (1 << 2) // 1010 OR 0100 = 1110
	fmt.Println(x)
	x &= ^(1 << 2) // 1110 AND 1011 = 1010
	fmt.Println(x)
	fmt.Println(SetBit(x,2,1))
}
func SetBit(x int64, i uint, value int) int64{
	if value == 1 {
		x |= (1 << i)
	}
	if value == 0 {
		x &= ^(1 << i)
	}
	return x
}
