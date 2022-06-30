package main

import (
	"fmt"
)

func main()  {

	var intChan chan int
	intChan = make(chan int, 3)

	//fmt.Printf("%v", intChan)

	intChan<- 10
	num := 211
	intChan<- num
	
	fmt.Printf("%v, %v", len(intChan), cap(intChan))

	num2 := <-intChan
	fmt.Println(num2)

	num3 := <-intChan
	fmt.Println(num3)

	fmt.Printf("%v, %v", len(intChan), cap(intChan))




}