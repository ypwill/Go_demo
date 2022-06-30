package main

import "fmt"

func main() {
	// map 声明
	var a = make(map[int]int)
	a[1] = 2
	a[2] = 3
	a[3] = 4
	fmt.Println(a)

	var city = map[string]string{
		"no1": "北京",
		"no2": "上海",
	}
	fmt.Println(city)

	fmt.Println("***********************")

	// map的 crud 操作

	b := make(map[string]int)
	b["小明"] = 99
	b["小红"] = 66
	b["小刚"] = 33
	fmt.Println(b)

	//delete(b, "小红")
	fmt.Println(b)

	// map 查找
	val, ok := b["小明1"]
	if ok {
		fmt.Printf("查询到了 %d", val)
	} else {
		fmt.Println("查询失败")
	}

	// map 遍历  使用 for range
	for k, val := range b {
		fmt.Printf("%v:%v\n", k, val)
	}

	fmt.Println(len(b))
}
