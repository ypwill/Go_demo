package main

import "fmt"

func main() {

	// 可以从控制台接收用户信息 （姓名、年纪、年薪、是否通过考试）
	// 方式一
	var name string
	var age int
	var salary float64
	var passed bool 
	// fmt.Println("请输入姓名") 
	// fmt.Scanln(&name)

	// fmt.Println("请输入年纪") 
	// fmt.Scanln(&age)

	// fmt.Println("请输入年薪") 
	// fmt.Scanln(&salary)

	// fmt.Println("请输入是否通过") 
	// fmt.Scanln(&passed)

	// fmt.Printf("姓名是%v,年纪%v,年薪%v, 是否通过考试%v", name, age, salary, passed)


	// 方式二
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &passed)
	fmt.Printf("姓名是%v,年纪%v,年薪%v, 是否通过考试%v", name, age, salary, passed)
	



	 


}