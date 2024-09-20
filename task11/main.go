package main

import "fmt"

func main() {
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})
	set1[1] = struct{}{}
	set1[2] = struct{}{}
	set1[16] = struct{}{}
	set1[5] = struct{}{}
	set1[8] = struct{}{}
	set1[9] = struct{}{}
	set1[13] = struct{}{}
	set2[7] = struct{}{}
	set2[2] = struct{}{}
	set2[54] = struct{}{}
	set2[5] = struct{}{}
	set2[4] = struct{}{}
	set2[9] = struct{}{}
	set2[13] = struct{}{}
	fmt.Println(set1, set2)
	fmt.Println(GetIntersection(set1, set2))
}

func GetIntersection(set1, set2 map[int]struct{}) map[int]struct{} {
	intersection := make(map[int]struct{})
	for k := range set1 {
		if _, found := set2[k]; found {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}
