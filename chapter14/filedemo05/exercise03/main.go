package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	
)

func main() {
	
	// 打开已经存在的文件  写5句 hello world
	filePath := "d:/a.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	str := "hello world\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}