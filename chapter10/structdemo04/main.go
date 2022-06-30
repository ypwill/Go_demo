package main

import "fmt"

type Stu struct {
	Name string
	Age int
}

func main() {

	stu1 := Stu{"小明", 18}
	fmt.Println(stu1)

	var stu2 = Stu{
		Name : "xiaohong",
		Age : 11,
	}
	fmt.Println(stu2)

	fmt.Println(stu1, stu2)

	// 方式2 返回结构体的指针类型
	var stu3 = &Stu{"小刚", 11}
	fmt.Println(&stu3)
	fmt.Println(stu3)

	stu4 := &Stu{
		Name : "小李",
		Age : 11,
	}
	fmt.Println(stu4)
	fmt.Println(*stu4)
	fmt.Println(&stu4)

}