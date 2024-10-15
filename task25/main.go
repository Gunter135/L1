package main

import (
	"fmt"
	"time"
)

// блокируем наш поток, ждем пока не пройдет duration
func Sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Sleeping for 2 seconds...")

	Sleep(2 * time.Second)

	fmt.Println("Woke up after 2 seconds!")

	go nonBlockingSleep(1, time.Millisecond*100)
	go nonBlockingSleep(2, time.Millisecond*400)
	go nonBlockingSleep(3, time.Millisecond*200)
	fmt.Println("Waiting another 2 seconds...")
	//Без WG просто ждем 2 секунды
	Sleep(2 * time.Second)
}

// если нам нужно для каждой горутины свой sleep то используем неблокирющий вариант, в котором горутина ждёт назависимо от других
func nonBlockingSleep(id int, duration time.Duration) {
	fmt.Printf("Goroutine %d: Sleeping for %v...\n", id, duration)
	<-time.After(duration)
	fmt.Printf("Goroutine %d: Woke up after %v\n", id, duration)
}
