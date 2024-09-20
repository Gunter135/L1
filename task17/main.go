package main

import "fmt"

func main() {

}

func binSearch(arr []int, key int) int {
	l := 1
	r := len(arr)
	var mid int
	for l <= r {
		mid = (l + r) / 2
		fmt.Println(l, mid, r)
		if arr[mid-1] == key {
			return mid - 1
		}
		if arr[mid-1] > key {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}