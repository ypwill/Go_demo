package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		fmt.Println("服务器在等在用户输入...\n", conn.RemoteAddr())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器端read error", err)
			return 
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println(err)
	}
	defer listen.Close()

	for {
		fmt.Println("服务器在等待客户端的连接...")
		conn, err := listen.Accept()	 
		if err != nil {
			fmt.Println("Accept()连接出错")
		} else {
			fmt.Printf("Accept() succ con = %v\n", conn.RemoteAddr())
		}
		go process(conn)
	}
}