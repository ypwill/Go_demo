package process

import (
	"fmt"
	"net"
	"go_code/chatroom/common/message"
	"go_code/chatroom/client/utils"
	"encoding/json"
	"encoding/binary"
	"os"
)

type UserProcess struct {

}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	// fmt.Printf("  userId = %d userPwd = %s\n", userId, userPwd)
	// return nil

	// 1.连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
	}

	defer conn.Close()

	// 2.通过conn 发送消息
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3. 创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 4.将LoginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return 
	}
	mes.Data = string(data)

	// 5. 将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return 
	}

	// 6. 接下来就是发送数据
	// 6.1 先把长度 发送给服务器   先获取data长度 -》表示长度的切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err = ", err)
		return 
	}

	  fmt.Printf("客户端,发送消息长度=%d 内容=%s", len(data), string(data))

	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err = ", err)
		return 
	}

	// 接收服务器返回的消息
	// 创建transfer
	tf := &utils.Transfer{
		Conn : conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err = ", err)
		return 
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {

		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		// 显示在线列表  遍历loginResMes.UsersId
		fmt.Println("当前在线用户列表如下:")
		for _, v := range loginResMes.UsersId {
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			//  初始化 客户端onlineUsers
			// user 信息构建
			user := &message.User{
				UserId : v,
				UserStatus : message.UserOnline,

			}
			onlineUsers[v] = user

		}
		fmt.Println()


		// 这里我们还需要起一个协程   该协程保持和服务端的通信  
		// 如果服务器有消息推送，则接收并显示在控制台
		go serverProcessMes(conn)

		// 进入成功后的菜单
		for {
			ShowMenu()
		}

	} else {
		fmt.Println(loginResMes.Error)
	}
	return 

}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {

	// 1.连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
	}

	defer conn.Close()

	// 2.通过conn 发送消息
	var mes message.Message
	mes.Type = message.RegisterMesType

	// 3. 创建一个LoginMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	// 4.将registerMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return 
	}
	mes.Data = string(data)

	// 5. 将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return 
	}

	// 6. 接下来就是发送数据
	// 6.1 先把长度 发送给服务器   先获取data长度 -》表示长度的切片

	tf := &utils.Transfer{
		Conn : conn,
	}
	// 发送data给服务器端
	err = tf.WritePkg(data) 
	if err != nil {
		fmt.Println("注册发送信息错误 err = ", err)
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err = ", err)
		return 
	}
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功！！！重新登录下吧~")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)

	}
	return 
}

