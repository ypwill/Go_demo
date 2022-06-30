package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 80000; i++ {
		intChan<- i 
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num , ok := <-intChan
		if !ok {
			break 
		}
		flag = true
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan<- num 
		}
	}
	fmt.Println("有一个 primeNum 协程因为取不到数据，退出")
	exitChan<- true 

}



func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	start := time.Now().Unix()

	go putNum(intChan)
	for i := 1; i <= 4; i++ {
		go primeNum(intChan, primeChan, exitChan) 
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		end := time.Now().Unix()
		fmt.Println("使用4个协程耗费的时间",end - start)

		close(primeChan)
	}()

	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("main 线程退出")



}