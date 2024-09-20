package main

import "fmt"

func main() {
	input := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	temperatures := make(map[int][]float32)
	for i := 0; i < len(input); i++ {
		temperatures[int(input[i]/10) * 10] = append(temperatures[int(input[i]/10) * 10], input[i])
	}
	fmt.Println(temperatures)
}