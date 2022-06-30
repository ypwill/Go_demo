package main

import (
	"fmt"
	"encoding/json"
) 

type Stu struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string
	Phone string
}


// 结构体序列化
func testStruct() {
	stu := Stu{
		Name : "小明",
		Age : 14,
		Gender : "男",
		Phone : "17746611780",
	}
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", data)
}


// map 序列化
func testMap() {
	a := make(map[string]interface{})
	a["name"] = "小红"
	a["Age"] = 11
	a["Phone"] = "1111111111"
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", data)
}

// 对切片进行序列化
func testSlice() {
	var a []map[string]interface{}
	m1 := make(map[string]interface{})
	m1["name"] = "小刚"
	m1["Age"] = 12
	m1["address"] = "北京朝阳"
	a = append(a, m1)

	m2 := make(map[string]interface{})
	m2["name"] = "小小"
	m2["Age"] = 11
	m2["address"] = "北京海淀"
	a = append(a, m2)

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", data)
}

func testInt() {
	i := 100
	data, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", data)

}



func main()  {
	
	testStruct()
	fmt.Println()
	testMap()

	fmt.Println()

	testSlice()
	fmt.Println()

	testInt()
	fmt.Println()

}