package main

import (
	"fmt"
	"go_code/chatroom/client/process"
)

var userId int
var userPwd string
var userName string

func main() {

	var key int
	//var loop = true

	for {
		fmt.Println("--------------------欢迎登陆多人聊天系统--------------------")
		fmt.Println("--------------------  1  登录聊天室")
		fmt.Println("--------------------  2  注册用户")
		fmt.Println("--------------------  3  退出系统")
		fmt.Println("--------------------  请选择(1-3):")

		fmt.Scanln(&key)
		switch key {
			case 1:
				fmt.Println("登录聊天室")
				fmt.Printf("请输入用户id:")
				fmt.Scanln(&userId)
				fmt.Printf("请输入用户密码:")
				fmt.Scanln(&userPwd)
				// 完成登陆
				up := &process.UserProcess{}
				up.Login(userId, userPwd)
			case 2:
				
				fmt.Println("注册用户")
				fmt.Printf("请输入用户id:")
				fmt.Scanln(&userId)
				fmt.Printf("请输入用户密码:")
				fmt.Scanln(&userPwd)
				fmt.Printf("请输入用户名字:")
				fmt.Scanln(&userName)
				up := &process.UserProcess{}
				up.Register(userId, userPwd, userName)
				//loop = false
			case 3:
				fmt.Println("退出系统")
				//loop = false
			default :
				fmt.Println("你的输入有误，请重新输入")
		}
	}
	
}

