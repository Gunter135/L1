package main

import (
	"fmt"
	"sync"
)

func main() {
	for i := 0; i < 100; i++ {
		array := []int{2, 4, 6, 8, 10}
		// array = channelSync(array)
		array = mutexSync(array)
		fmt.Println(array)
	}
}

func channelSync(array []int) []int {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	for i := 0; i < len(array); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-ch
			array[i] = array[i] * array[i]
			ch <- struct{}{}
		}(i)
	}
	wg.Wait()
	// time.Sleep(1000)
	return array
}

func mutexSync(array []int) []int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < len(array); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			array[i] = array[i] * array[i]
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	// time.Sleep(1000)
	return array
}
