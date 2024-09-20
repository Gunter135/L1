package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	done := make(chan bool)
	workChan := make(chan interface{})
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 0 {
		log.Fatal("Time must be a number or greater than 0")
	}
	timer := time.NewTimer(time.Second * time.Duration(n))
	go func() {
		for {
			select{
			case <-timer.C:
				done <- true
				close(workChan)
				return
			default:
				workChan <- struct{}{}
			}
		}
	}()
	go func() {
		for {
			_,status := <-workChan
			if !status {
				return
			}
		}
	}()
	for {
		select {
		case <-done:
			fmt.Println("Finished ✓")
			return
		default:
			fmt.Println("Waiting...")
		}
	}
}
