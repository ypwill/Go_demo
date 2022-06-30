package main

import (
	"fmt"
	"strconv"
	"time"
)


func test() {
	for i:= 1; i <= 100; i++ {
		fmt.Println("test() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main()  {

	go test() // 开启了一个协程     主线程执行完  协程也会结束
	for i:= 1; i <= 5; i++ {
		fmt.Println("main() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}