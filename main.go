package main

import (
	"fmt"
	"os"
)

func main() {
	//Create a folder/directory at a full qualified path

	//Create a folder/directory at a full qualified path
	for i := 1; i < 27; i++ {
		folder :=fmt.Sprintf("%s%d","task",i);
		err := os.Mkdir(folder, 0755)
		if err != nil {
			fmt.Println(err)
		}
	}
	
}