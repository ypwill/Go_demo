package main

import(
	 "fmt"
	 "sort"
)

func main() {
	// monster := make(map[string]int)
	// monster["小怪1"] = 1
	// fmt.Println(monster) 

	// map 切片
	monster := make([]map[string]int, 2)
	if monster[0] == nil {
		monster[0] = make(map[string]int, 2)
		monster[0]["num01"] = 10 
		monster[0]["num02"] = 1000
	}
	fmt.Println(monster)


	nums := map[int]int{
		100 : 1,
		21 : 2,
		32 : 3,
		44 : 4,
	}
	fmt.Println(nums)

	// map 排序 输出
	// 先将map 放入到一个切片  对切片排序  输出	
	var key []int
	for k, _ := range nums {
		key = append(key, k)
	}
	// 对 key排序
	sort.Ints(key)
	fmt.Println(key)

	for i := 0; i < len(key); i++ {
		fmt.Printf("%d ", nums[key[i]])
	}



}