package main

import "fmt"

func main() {

}
func partition(l int, h int, arr **[]int) int {
	ptr := **arr
	pivot := l
	for l < h {
		fmt.Println(l)
		for ptr[l] <= ptr[pivot] {
			l++
		}
		for ptr[h] > ptr[pivot] {
			h--
		}
		if l < h {
			ptr[l], ptr[h] = ptr[h], ptr[l]
		}
	}

	ptr[pivot], ptr[h] = ptr[h], ptr[pivot]

	return h
}

func quickSort(l int, h int, arr **[]int) {

	if l < h {
		mid := partition(l, h, arr)
		quickSort(l, mid, arr)
		quickSort(mid+1, h, arr)
	}
}