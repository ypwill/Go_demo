package main

import "fmt"

func main() {
	var arr = [2][2]int{{1, 2,}, {2, 3}}
	fmt.Println(arr)

	// for i := 0; i < len(arr); i++ {
	// 	for j := 0; j < len(arr[0]); j++ {
	// 		//arr[i][j] = 1
	// 		fmt.Printf("%d ", arr[i][j])
	// 	}
	// 	fmt.Println()
	// }
	for i, _ := range arr {
		for j, _ := range arr[i] {
			fmt.Printf("%d ", arr[i][j])
		} 
		fmt.Println()
	}
	
}