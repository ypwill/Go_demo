package main

import "fmt"
func main() {
	
	// longday := 97
	// var week int
	// var day int
	// week = longday / 7
	// day = longday % 7
	// fmt.Printf("总共有 %d 个week零%dday", week, day)

	// var a int = 3
	// var b int = 4
	// fmt.Println(a > b)
	// fmt.Println(a == b)
	// fmt.Println(a != b)
	// flag := a > b
	// fmt.Println(flag)

	var age = 40
	if age >= 30 && age <= 50 {
		fmt.Println("yes")
	}
	if age >= 30 && age < 40 {
		fmt.Println("yes")
	}

	a := 5
	b := 8
	a = a - b
	b = a + b
	a = b - a
	fmt.Printf("%d %d", a, b)
	

}