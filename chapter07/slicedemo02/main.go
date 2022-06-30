package main

import (
	"fmt"
)

func main() {

	// 切片的基本使用
	var intArr [5]int = [5]int{1, 2, 3, 4, 5}
	// 声明定义 一个切片
	// intArr[1:3] 表示slice 引用到intArr数组下标从 [1-3)  不包括3
	slice := intArr[1:3]
	fmt.Println(intArr)
	fmt.Println(slice)
	fmt.Println(len(slice))	 // 元素个数
	fmt.Println(cap(slice))  // 容量 可以动态变化

	fmt.Println("**************************") 

	// slice 的三种使用方式
	// 1.定义一个切片 去引用之前已经定义的数组
	var array = [5]int{11,22,33,44,55}
	arraySlice := array[1:3]
	fmt.Println(arraySlice)

	// 2.使用make  参数 1.切片类型 2.切片大小 3.容量
	var slice02 = make([]int, 4, 10)
	fmt.Println(slice02)
	fmt.Printf("%v %v", len(slice02), cap(slice02))
	slice02[3] = 1
	fmt.Println(slice02[3])

	// 3.定义一个切片，直接指定具体数组
	var slice03 []int = []int{1, 2, 3}
	fmt.Println(slice03)
	fmt.Println(len(slice03))
	fmt.Println(cap(slice03))

	fmt.Println("************************")

	for i := 0; i < len(slice03); i++ {
		fmt.Println(slice03[i])
	}

	for _, val := range slice03 {
		fmt.Println(val)
	}

	fmt.Println("************************")

	// append函数
	var slice04 = []int{1, 2, 3, 4, 5}
	slice04 = append(slice04, slice04...)
	slice05 := append(slice04, 6, 7)

	fmt.Println(slice04)
	fmt.Println(slice05)
	



}