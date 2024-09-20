package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	counter := &CounterA{}
	// counter := &CounterM{}
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				counter.Increment()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("The counter is: %d\n", counter.GetCounter())
}

type CounterM struct {
	m       sync.Mutex
	counter int
}

func (c *CounterM) Increment() {
	c.m.Lock()
	defer c.m.Unlock()
	c.counter++
}

func (c *CounterM) GetCounter() int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.counter
}

type CounterA struct {
	counter atomic.Int32
}

func (c *CounterA) Increment() {
	c.counter.Add(1)
}

func (c *CounterA) GetCounter() int {
	return int(c.counter.Load())
}
