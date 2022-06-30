package main

import (
	"fmt"
	"encoding/json"
)

type Stu struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func (stu Stu) test() {
	fmt.Println(stu.Name)
}

type Point struct {
	x int
	y int
}

type Rect struct {
	left, right Point
}


func main() {
	// 创建结构体变量的几种方法
	// 1. 直接创建
	var stu1 Stu
	fmt.Println(stu1)

	// 2. 赋值创建
	stu2 := Stu{"小明", 18}
	fmt.Println(stu2)

	// 3. &  
	// var stu3 *Stu = new(Stu)
	//(*stu3).name = "xiaohong"
	// stu3.name = "xiaohong"
	// (*stu3).age = 11
	// fmt.Println(*stu3)

	// 4.
	var stu4 *Stu = &Stu{}
	fmt.Println(stu4)

	fmt.Println("**********************")

	r1 := Rect{Point{1, 2}, Point{3, 4}}
	r2 := Rect{Point{1, 2}, Point{5, 6}}
	fmt.Println(r1)
	fmt.Println(r2)

	fmt.Println("**********************")

	// 序列化使用 转json
	stu5 := Stu{"bob", 18}
	jsonStu, err := json.Marshal(stu5)
	if err != nil {
		fmt.Println("出现错误")
	}
	fmt.Println(stu5)
	fmt.Println(string(jsonStu))   // 结构体 属性名 需要大写

	fmt.Println("**********************")

	stu6 := Stu{"小小鹏", 18}
	stu6.test()

	 



}