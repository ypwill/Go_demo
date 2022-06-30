package process

import (
	"fmt"
	"os"
	"net"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
	"encoding/json"
)

// 显示登陆成功后的界面...
func ShowMenu() {

	fmt.Println("---------------恭喜xxx登陆成功---------------")
	fmt.Println("---------------1. 显示在线用户列表---------------")
	fmt.Println("---------------2. 发送消息---------------")
	fmt.Println("---------------3. 信息列表---------------")
	fmt.Println("---------------4. 退出系统---------------")
	fmt.Printf("请选择(1-4):")
	var key int
	var content string
	// 因为我们总会用到smsProcess  所以放在外面
	smsProcess := &SmsProcess{}
	fmt.Scanln(&key)
	switch key {
		case 1 :
			//fmt.Println("显示在线用户列表")
			outputOnlineUser()
		case 2 :
			fmt.Println("请输入你想对大家说点什么:")
			fmt.Scanln(&content)
			smsProcess.SendGroupMes(content)
		case 3 :
			fmt.Println("信息列表")
		case 4 :
			fmt.Println("你选择退出系统")
			os.Exit(0)
		default :
			fmt.Println("你输入的有问题，请重新输入...")
	}
}
 
func serverProcessMes(conn net.Conn) {

	tf := &utils.Transfer{
		Conn : conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err = ", err)
			return 
		}
		fmt.Println(mes)
		// 如果读取到了数据 又是下一步处理逻辑
		// 根据消息类型进行下一步解析
		switch mes.Type {
			case message.NotifyUserStatusMesType :
				// 处理 1. 取出 NotifyUserStatusMes  保存用户状态  保存在客户 map[int]User 
				var notifyUserStatusMes *message.NotifyUserStatusMes
				json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
				updateUserStatus(notifyUserStatusMes)
			case message.SmsMesType :
				outputGroupMes(&mes)
			default :
				fmt.Println("服务端返回了未知的消息类型")
		}
		
	}
}