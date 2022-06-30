package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main() {
	
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("连接出错")
	}
	// 功能一：客户端可以发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin) // os.Stdio代表标准输入 从控制台
	if err != nil {
		fmt.Println("创建reader失败")
	}
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取失败")
		}
		str = strings.Trim(str, " \r\n")
		if str == "exit" {
			break
		}
		
		_, err = conn.Write([]byte(str + "\n"))
		if err != nil {
			fmt.Println("conn.Write失败")
		}
	}
	



}