package main

import (
	"fmt"
	"runtime"
)

func main()  {

	i := runtime.NumCPU()
	fmt.Println(i)
	runtime.GOMAXPROCS(i - 1)

}