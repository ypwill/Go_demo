package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)

func CopyFile(destFileName string, sourceFileName string) (written int64, err error) {
	srcFile, err := os.Open(sourceFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	destFile, err := os.OpenFile(destFileName, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(destFile)
	defer destFile.Close()
	return io.Copy(writer, reader)
}

func main() {

	sourceFile := "d:/1.jpg"
	destFile := "d:/2.jpg"
	_, err := CopyFile(destFile, sourceFile)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Println("粗错啦")
	}

}