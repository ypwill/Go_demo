package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4 ,5}
	var slice = make([]int, 1)
	copy(slice, a)
	fmt.Println(slice)
}
