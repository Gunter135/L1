package main

import "fmt"

func main() {
	GetType(1)
}

func GetType(input interface{}) {
	switch v := input.(type) {
	case int:
		fmt.Printf("It's an integer! %d\n",v)
	case string:
		fmt.Printf("It's a string! %s\n", v)
	case bool:
		fmt.Printf("It's a boolean! %t\n",v)
	case chan int:
		fmt.Printf("It's an integer channel!\n")
	case chan string:
		fmt.Printf("It's a string channel!\n")
	case chan bool:
		fmt.Printf("It's a bool channel!\n")
	default:
		fmt.Println("Unknown type")
	}
}