package main

import "fmt"

func test() {
	defer func(){
		error := recover()
		if error != nil {
			fmt.Println(error)
		}
	}() 
	n1 := 100
	n2 := 0
	fmt.Println(n1 / n2)

}

func main() {

	test()
	fmt.Println("hello")
}