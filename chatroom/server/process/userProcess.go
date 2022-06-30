package processes

import (
	"fmt"
	"net"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
	"go_code/chatroom/server/model"
	"encoding/json"
)

type UserProcess struct {
	Conn net.Conn
	UserId int
}

// 编写通知所有在线用户的方法
func (this *UserProcess) NotifyOthersOnlineUsers(userId int){
	// 需要遍历 userMgr.onlineUsers 
	for id, up := range userMgr.onlineUsers {
		// 过滤掉自己
		if id == userId {
			continue
		}
		// 开始通知 （写一个方法）
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {
	// NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status =  message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal fail err = ", err)
		return 
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal fail err = ", err)
		return 
	}
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err = ", err)
		return 
	}
}



// 编写一个 ServerProcessLogin 函数  专门处理登录逻辑
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err!= nil {
		fmt.Println("json.Unmarshal fail err = ", err)
		return 
	}

	// 声明 resMes    构建message  结构体并返回
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes

	// 此时需要到redis里去查询
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
	 		loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
	 		loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
	 		loginResMes.Error = "服务器内部错误..."
		}
	} else {
		loginResMes.Code = 200

		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		this.NotifyOthersOnlineUsers(loginMes.UserId)

		// 将当前在线用户的id 放入loginResMes 的 UserIds
		// 遍历userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}

		fmt.Println("登陆成功", user)
	}


	// if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	// 	// 合法
	// 	loginResMes.Code = 200
	// } else {
	// 	// 不合法
	// 	loginResMes.Code = 500
	// 	loginResMes.Error = "该用户不存在，请注册后再使用..."
	// }


	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail err = ", err)
		return 
	}
	// 将序列化后的data 赋值给 构建的 message   然后再序列化
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail err = ", err)
		return 
	}
	// 发送 data  我们将其封装到 writePkg函数
	// 因为使用分层（mvc） 先创建一个Transer
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return 
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err!= nil {
		fmt.Println("json.Unmarshal fail err = ", err)
		return 
	}

	// 声明 resMes    构建message  结构体并返回
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}


	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail err = ", err)
		return 
	}
	// 将序列化后的data 赋值给 构建的 message   然后再序列化
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail err = ", err)
		return 
	}
	// 发送 data  我们将其封装到 writePkg函数
	// 因为使用分层（mvc） 先创建一个Transer
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return 
}
