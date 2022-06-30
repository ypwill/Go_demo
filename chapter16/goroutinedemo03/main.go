package main

import (
	"fmt"
	"time"
	"sync"
)

var (
	m = make(map[int]int)
	lock sync.Mutex
)

func test (n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i 
	}
	lock.Lock()
	m[n] =res
	lock.Unlock()
}

func main()  {

	for i := 1; i <= 200; i++ {
		go test(i)
	}
	//  主线程执行完成之后，协程就会结束  所以需要让主线程等待一段时间
	time.Sleep(time.Second)  // 主线程等待协程完成时间不确定     
	lock.Lock()
	for i, v := range m {
		fmt.Printf("m[%d] = %d\n", i, v)
	}
	lock.Unlock()



}