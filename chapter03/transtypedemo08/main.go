package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
	"strconv"
)

func main() {
	
	// var i int32 = 100
	// var n1 float32 = float32 (i) 
	// fmt.Println(n1)


	// 基本数据类型与 string类型的相互转换
	num1 := 99
	num2 := 23.456
	b := false
	myChar := 'a'
	var str string 

	// 使用第一种方式转换
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str %T, str = %q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str %T, str = %q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str %T, str = %q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str %T, str = %q\n", str, str)


	// 方式二  使用 strconv 包函数
	num3 := 99
	num4 := 23.456
	c := false

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str %T, str = %q\n", str, str)


	// 说明 strconv.FormatFloat(num4, 'f', 10, 64)
	// f 代表一种格式（数据转化的格式）  10 小数精度  64 表示小数是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str %T, str = %q\n", str, str)


	str = strconv.FormatBool(c)
	fmt.Printf("str %T, str = %q\n", str, str)

	num5 := 600
	str = strconv.Itoa(num5)
	fmt.Printf("str %T, str = %q\n", str, str)



}