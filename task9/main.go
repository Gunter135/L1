package main

import (
	"fmt"
	"sync"
)

func main() {
	// input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	input := generateArray(1000)
	inputCh := make(chan int)
	outputCh := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(inputCh, outputCh, &wg)
	go printer(outputCh, &wg)
	go func() {
		for _, val := range input {
			inputCh <- val
		}
		close(inputCh)
	}()
	wg.Wait()

}

func worker(inputCh <-chan int, outputCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range inputCh {
		result := val * 2
		outputCh <- result
	}
	close(outputCh)
}
func printer(inputCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range inputCh {
		fmt.Println(val)
	}
}

func generateArray(size int) []int {
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = i + 1
	}
	return array
}
