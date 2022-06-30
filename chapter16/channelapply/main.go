package main

import (
	"fmt"
	"time"
 ) 

func WriteData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan<- i 
		fmt.Println("放入", i)
		time.Sleep(time.Second)
	}
	close(intChan)
}

func ReadChan(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <- intChan
		if !ok {
			break
		} 
		fmt.Println("读取", v)
		time.Sleep(time.Second)
	}
	exitChan<- true
	close(exitChan)
}


func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go WriteData(intChan)
	go ReadChan(intChan, exitChan)

	for {
		_, ok := <- exitChan
		if !ok {
			break
		}
	}

}