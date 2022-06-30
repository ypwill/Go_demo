package main

import (
	"fmt" 
)

// 演示golang中指针类型
func main() {
	
	// var i = 10
	// fmt.Println(&i)

	// var ptr *int = &i
	// fmt.Println(ptr)
	// fmt.Println(&ptr)
	// fmt.Println(*ptr)

	var i = 100
	fmt.Println(&i)

	var ptr *int = &i
	fmt.Println(ptr)

	*ptr = 200
	fmt.Println(i)
	

}