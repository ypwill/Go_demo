package main

import (
	"fmt"
	"go_code/chapter10/factory/model"
)

func main() {

	// 使用工厂模式 解决不能访问 小写结构体
	// stu1 := model.student{"小明", 18}
	// fmt.Println(stu1)

	// 定义student 结构体是首字母小写，我们可以通过工厂模式解决
	stu2 := model.NewStudent("小明", 18)
	fmt.Println(stu2.Name)
	// fmt.Println(stu2.age)
	fmt.Println(stu2.GetAge())



}