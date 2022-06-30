package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)

func main() {
	file, err := os.Open("d:/a.txt")
	if err != nil {
		fmt.Println("open file error = ", err)
	}

	defer file.Close()

	read := bufio.NewReader(file)
	// 循环读取
	for {
		str, err := read.ReadString('\n') // read跟file绑定  每次读取到换行符就停止
		if err == io.EOF {   // io.EOF 表示文件末尾
			break
		}
		fmt.Print(str)
		
	}
	fmt.Println("文件读取结束")
}