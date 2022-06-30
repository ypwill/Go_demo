package main

import (
	"fmt"
	"reflect"
 ) 


func reflectTest01(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Println(rTyp)

	rVal := reflect.ValueOf(b)

	n2 := 2 + rVal.Int() // 获取 value 类型的 rVal 的值
	fmt.Println(n2)

	fmt.Printf("%T", rVal)

	

	iV := rVal.Interface()
	num := iV.(int)
	num++
	fmt.Println(num)

	fmt.Println("*****************")
}

func reflectTest02(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Println(rTyp)

	rVal := reflect.ValueOf(b)
	fmt.Println(rVal)
	fmt.Printf("%T", rVal)


	iV := rVal.Interface()
	stu := iV.(Stu)
	fmt.Println(stu.Name)

	fmt.Println(rVal.Kind())

}

type Stu struct {
	Name string
	Age int
}



 func main()  {
	var num int = 100
	reflectTest01(num)


	stu := Stu{"小明", 12}
	reflectTest02(stu)

 }