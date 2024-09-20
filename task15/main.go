package main

import "fmt"

// Так делать не надо, в хипе будет лежать большая стринга, в стеке лежит ссылка на маленький кусочек -> gc не чистит -> memory leak
// не совсем лучшая идея делать глобальную переменную в данном случае
// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

func main() {
	justString := someFunc()
	fmt.Println(justString)
	// fmt.Println(&justString)
}
func someFunc() string {
	v := createHugeString(1 << 10)
	if len(v) > 100 {
		return string(v[:100])
	}
	return string(v)
}

func createHugeString(size int) []byte {
	letters := "abcdefg"
	ls := len(letters)
	result := make([]byte, size)
	for i := 0; i < size; i++ {
		result[i] = letters[i%ls]
	}
	return result
}
