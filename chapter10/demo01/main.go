package main

import "fmt"


type Cat struct {
	name string
	age int
	color string
}

func detail(cats map[string]Cat, name string) {
	flag := false
	for k, v := range cats {
		if (name == k){
			flag = true
			fmt.Println(v)
		}
	}
	if (!flag){
		fmt.Println("张老太没有这只猫猫")
	}
}

func main() {

	cats := make(map[string]Cat)
	cat1 := Cat{"小白", 3, "white"}
	cat2 := Cat{"小花", 100, "flower"}
	cats["小白"] = cat1
	cats["小花"] = cat2
	fmt.Println(cats)

	detail(cats, "小")
}