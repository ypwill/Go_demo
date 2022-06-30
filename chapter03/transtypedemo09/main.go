package main

import (
	"fmt" 
	"strconv"
)

func main() {
	
	// string 类型转基本数据类型
	str := "1234"
	var num1 int64
	num1, _ = strconv.ParseInt(str, 10, 64)
	fmt.Printf("num1 = %d   num1 type %T\n", num1, num1)

	var str1 string = "true"
	var b bool
	b, _ = strconv.ParseBool(str1)
	fmt.Printf("b type %T   %v\n", b, b)


	var str2 = "12.3456"
	var num2 float64
	num2, _ = strconv.ParseFloat(str2, 64)
	fmt.Printf("num2 = %f   num2 type %T\n", num2, num2)





}