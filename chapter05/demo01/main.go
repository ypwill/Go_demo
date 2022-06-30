package main

import "fmt"
func main() {
	
	// var score int
	// fmt.Scanln(&score)

	// switch  {
	// 	case score > 60 :
	// 		fmt.Println("合格")
	// 	case score < 60 :
	// 		fmt.Println("不合格")
	// 	default:
	// 		fmt.Println("成绩输入有误")
	// }

	// for i := 0; i < score; i++ {
	// 	fmt.Println("hello world")
	// }

	// str := "hello,world"
	// for index, val := range str {
	// 	fmt.Printf("%d %c\n",index, val)
	// }
	i := 5
	for a := 1; a <= i; a++ {
		for b := 1; b <= i - a; b++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= a * 2 - 1; j++ {
			if j == 1 || j == a * 2 - 1 || a == i {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
			
		}
		fmt.Println()
	}
	 
	

}