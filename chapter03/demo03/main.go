package main

import "fmt"

var n1 = 100
var n2 = 100
var name = "jack"

var (
	n3 = 100
	n4 = 100
	name1 = "bob"
)

func main() {
	// golang 一次性声明多个变量
	// var n1, n2, n3 int
	// fmt.Println("n1 = ", n1, "n2 = ", n2, "n3 = ", n3)

	// 一次性声明多个变量 
	// var n1, name, n3 = 100, "tmo", 888 
	// fmt.Println("n1 = ", n1, "name = ", name, "n3 = ", n3)

	// 可以使用类型推导
	n1, name, n3 := 100, "tmo", 888
	fmt.Println("n1 = ", n1, "name = ", name, "n3 = ", n3)
	 


}