package main

import (
	"fmt"
)

func BubbleSort(arr *[10]int) {
	fmt.Println(*arr)
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - i - 1; j++ {
			if (arr[j] > arr[j + 1]) {
				tmp := arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = tmp
			}
		}
	}
}

func Find (arr *[10]int, target int)  int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func Find02(arr *[10]int, target int) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := (l + r) / 2
		fmt.Printf("mid = %d  ", mid)
		if arr[mid] == target {
			return mid
		} else if (arr[mid] > target) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}



func main() {
	arr := [10]int{9, 3, 4, 6, 2, 8, 5, 1, 0, 7}
	
	// BubbleSort(&arr)
	// fmt.Println(arr)

	fmt.Println("******************************")

	// 顺序查找

	pos := Find(&arr, 8)

	fmt.Println(pos)

	// 二分查找
	BubbleSort(&arr)
	pos1 := Find02(&arr, 3)
	fmt.Println(pos1)

	




}