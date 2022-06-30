package main

import "fmt"

func main() {
	str := "hello@world"
	// 使用切片处理 取出 world
	slice := str[6:]
	fmt.Println(slice)

	// string 不可修改  想要修改可以转为[]byte 或[]rune 修改然后再转为string
	// 将 hello@world 改为 aello@world
	// arr1 := []byte(str)
	// arr1[0] = 'a'
	// str = string(arr1)
	// fmt.Println(str)

	// 要是含有中文  转为[]byte数组后，可以处理数字和英文，但是不能处理中文  byte按字节来处理，一个中文三个字节
	// 解决办法 转为 []rune ，因为rune是按字符来处理的 
	arr1 := []rune(str)
	arr1[0] = '北'
	str = string(arr1)
	fmt.Println(str)


	
}
