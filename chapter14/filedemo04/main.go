package main

import (
	"fmt"
	"bufio"
	"os"
	
)

func main() {
	
	// 创建一个新文件  写5句 hello world
	filePath := "d:/b.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	str := "hello go\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	writer.Flush()


	
}