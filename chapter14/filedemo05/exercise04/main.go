package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "io"
	"io/ioutil"
	
)

func main() {
	
	// 先读取文件内容
	// sourcePath := "d:/a.txt"
	// sourcefile, err := os.OpenFile(sourcePath, os.O_RDONLY, 0666)

	// reader := bufio.NewReader(sourcefile)

	// destPath := "d:/b.txt"
	// destFile, err := os.OpenFile(destPath, os.O_WRONLY | os.O_TRUNC, 0666)

	// writer := bufio.NewWriter(destFile)

	// var str string
	// for {
	// 	str, err = reader.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	writer.WriteString(str)
	// }
	// writer.Flush()

	// 方式二

	sourcePath := "d:/a.txt"
	destPath := "d:/b.txt"
	content, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(destPath, content, 0666)
	if err != nil {
		fmt.Println(err)
	}






}