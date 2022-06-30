package main

import "fmt"

type Point struct {
	x int
	y int
}


func main() {

	// 类型断言  
	var a interface{}
	point := Point{1, 2}
	a = point
	var b Point

	// 把 b直接赋给a是不可以的
	// b = a   这样是不行的
	b = a.(Point)
	fmt.Println(b)

	// 一般数据类型的断言  
	// var x interface{}
	// c := 3.14
	// x = c
	// y := x.(float64)
	// fmt.Println(y)

	// 带错误检测的类型断言 
	var x interface{}
	c := 3.14
	x = c
	y, flag := x.(float64)
	if (!flag){
		fmt.Println("类型转换失败")
	} else {
		fmt.Println(y)
	}
	fmt.Println("OKOKOKO")
}