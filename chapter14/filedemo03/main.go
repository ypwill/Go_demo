package main

import (
	"fmt"
	"io/ioutil"
	
)

func main() {
	file := "d:/a.txt"
	content, _ := ioutil.ReadFile(file)
	fmt.Printf("%s", content)
	fmt.Printf("%v", string(content))

	
}