package main

import (
	"fmt"
	_"bufio"
	_"os"
	
)

func main() {
	
	// 打开已经存在的文件  写5句 hello world
	// filePath := "d:/a.txt"
	// file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND, 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// defer file.Close()

	// str := "hello world\n"
	// writer := bufio.NewWriter(file)
	// for i := 0; i < 5; i++ {
	// 	writer.WriteString(str)
	// }

	// writer.Flush()
	a := 1
	b := 2
	a, b = b, a 
	fmt.Println(a, b)
}