package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func stopDoneChannel() {
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	done <- true
}

func stopContext() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
}

func stopChannelClose() {
	stop := make(chan struct{})

	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	close(stop)
}

func stopWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine running (waitgroup)...")
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("Goroutine stopped by WaitGroup.")
	}()

	wg.Wait()
}

func main() {
	stopDoneChannel()
	stopContext()
	stopChannelClose()
	stopWaitGroup()
}
