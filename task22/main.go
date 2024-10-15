package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("1048577", 10) // 2^20 + 1
	b.SetString("2097153", 10) // 2^21 + 1

	fmt.Println(sum(*a, *b))
	fmt.Println(sub(*a, *b))
	fmt.Println(mul(*a, *b))
	fmt.Println(div(*a, *b))
}

func sum(a big.Int, b big.Int) *big.Int {
	sum := new(big.Int).Add(&a, &b)
	return sum
}
func sub(a big.Int, b big.Int) *big.Int {
	sub := new(big.Int).Sub(&a, &b)
	return sub
}
func div(a big.Int, b big.Int) *big.Int {
	div := new(big.Int).Div(&a, &b)
	return div
}
func mul(a big.Int, b big.Int) *big.Int {
	mul := new(big.Int).Mul(&a, &b)
	return mul
}
