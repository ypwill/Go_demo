package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test(arr [4]int){
	arr[0] = -1
	fmt.Println(arr[0])
}
func test02(arr *[4]int){
	(*arr)[1] = -1
	fmt.Println(arr[1])
}
func main() {

	// 数组
	// var array [6]int
	// for i := 0; i < 6; i++ {
	// 	fmt.Scanln(&array[i])
	// }
	// fmt.Println(array)

	// 数组初始化方式
	var arr1 [4]int = [4]int {1,2,3,4}
	fmt.Println(arr1)

	var arr2 = [4]int {1,2,3,4}
	fmt.Println(arr2)

	var arr3 = [...]int {1,2,3,4}
	fmt.Println(arr3)

	var arr4 = [4]int {1:800, 2:200}
	fmt.Println(arr4)

	// 数组遍历方式  常规  和 for range
	for _, val := range arr4 {
		fmt.Println(val)
	}
	test(arr4)
	fmt.Println(arr4)

	fmt.Println("___________________________________")

	test02(&arr4)
	fmt.Println(arr4)

	fmt.Println("___________________________________")

	var arr5 [26]byte
	arr5[0] = 'A'
	for i := 1; i < 26; i++ {
		arr5[i] = arr5[i - 1] + 1
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", arr5[i])
	}

	fmt.Println("___________________________________")



	rand.Seed(time.Now().Unix())
	var arr6 [5]int
	for i := 0; i < len(arr6); i++ {
		arr6[i] = rand.Intn(100)
	}
	fmt.Println(arr6)
	for i := len(arr6) - 1; i >= 0; i-- {
		fmt.Println(arr6[i])
	}





}