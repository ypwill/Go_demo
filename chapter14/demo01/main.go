package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("d:/a.txt")
	if err != nil {
		fmt.Println("open file error = ", err)

	}
	fmt.Printf("%v", file)
	
	err = file.Close()
	if err != nil {
		fmt.Printf("%v", err)
	}
	
}