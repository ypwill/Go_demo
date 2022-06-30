package main

import (
	"fmt"
	"encoding/json"
) 

type Stu struct {
	Name string 
	Age int
	Gender string
	Phone string
}
// 
func UnSerialStruct() {

	str := "{\"Name\":\"小明\",\"Age\":14,\"Gender\":\"男\",\"Phone\":\"17746611780\"}"

	var stu Stu
	err := json.Unmarshal([]byte(str), &stu)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stu)
}

func UnmarshalMap() {
	str := "{\"Age\":12,\"address\":\"北京朝阳\",\"name\":\"小刚\"}"
	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)

}

func UnmarshalSlice() {
	str := "[{\"Age\":12,\"address\":\"北京朝阳\",\"name\":\"小刚\"},{\"Age\":11,\"address\":\"北京海淀\",\"name\":\"小小\"}]"
	var a []map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)

}


func main() {

	UnSerialStruct()
	fmt.Println()

	UnmarshalMap()
	fmt.Println()

	UnmarshalSlice()


}
