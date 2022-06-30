package main

import (
	"fmt"
	_"go_code/chapter06/demo01/utils"
)
func main() {

	// var i int = 1
	// var j int = 2
	// var res int
	// res = utils.Sum(i, j)
	// fmt.Println(res)


	// 匿名函数
	res := func (n1 int, n2 int) int {
		sum := n1 + n2
		return sum
	}(1, 2)
	fmt.Println(res)


	a := func (n1 int, n2 int) int {
		sum := n1 * n2
		return sum
	}

	res1 := a(3, 4)
	fmt.Println(res1)

	



	




}
