package process

import (
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/client/model"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

var CurUser model.CurUser  // 在登陆完成之后 完成对 curUser的初始化

// 显示当前在线用户
func outputOnlineUser(){
	fmt.Println("当前在线用户列表:")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

// 编写一个方法处理 返回的 NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId : notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}

