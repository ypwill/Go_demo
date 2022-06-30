package main

import "fmt"

func main()  {
	// var intChan chan int
	// intChan = make(chan int, 10)
	intChan := make(chan int, 10)
	intChan<- 10
	intChan<- 20
	close(intChan)

	intChan02 := make(chan int, 100)
	for i:= 0; i < 100; i++ {
		intChan02<- i + 1
	}

	// for i := 0; i < len(intChan02); i++ {   // len(intChan02)长度发生变化
	// 	fmt.Println(<-intChan02)
	// }
	// 遍历没有关闭管道 会出现 deadlock 错误  
	close(intChan02)
	for v := range intChan02 {
		fmt.Println(v)
	}


}