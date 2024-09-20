package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	workers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("The amount of workers must be a number")
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	taskChan := make(chan interface{})
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(taskChan, &wg)
	}
	Loop:
		for {
			select {
			case <-c:
				close(taskChan)
				break Loop
			default:
				taskChan <- struct{}{}
			}
		}
	fmt.Println("Terminating process...")
	wg.Wait()
}

func worker(workCh <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range workCh {
		//bottleneck
		fmt.Printf("Input data: %d\n", data)
	}
}
func taskGenerator(taskChan chan<- interface{}, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := i; j < i+1000; j++ {
		taskChan <- struct{}{}
	}
}