package main

import "fmt"


type Stu struct {
	name string
	age int
	address string
}

func main() {

	stuArr := make(map[string]Stu)
	stu1 := Stu{"小明", 12, "北京"}
	stu2 := Stu{"小红", 10, "南京"}
	stu3 := Stu{"小刚", 11, "西京"}
	stuArr["小明"] = stu1 
	stuArr["小红"] = stu2 
	stuArr["小刚"] = stu3
	fmt.Println(stuArr) 
}