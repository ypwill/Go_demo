package main

import (
	"fmt"
	"sort"
	"math/rand"
)

// 1. 声明hero结构体
type Hero struct {
	Name string
	Age int
}
// 2. 声明 一个 hero结构体切片类型
type HeroSlice []Hero

// 3. 实现 Sort 接口的方法
func (hs HeroSlice) Len() int{
	return len(hs)
}

// Less决定是使用升序还是降序排序     我们这是使用按年龄升序排列
func (hs HeroSlice) Less(i ,j int) bool{
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i ,j int){
	tmp := hs[i]
	hs[i] = hs[j]
	hs[j] = tmp
}


func main()  {
	var intSlice = []int{1, -3, 5, 3, 9, 6}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	// 对一个结构体切片排序
	// 使用自带的Sort 方法   实现data 接口中的  3个方法
	var heroSlice HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name : fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age : rand.Intn(100),
		}
		heroSlice = append(heroSlice, hero)
	}
	for _, v := range heroSlice {
		fmt.Println(v)
	}

	fmt.Println("***********************")

	sort.Sort(heroSlice)
	for _, v := range heroSlice {
		fmt.Println(v)
	}

	 
}