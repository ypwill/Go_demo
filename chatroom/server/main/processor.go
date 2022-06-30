package main

import (
	"fmt"
	"net"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/process"
	"go_code/chatroom/server/utils"
	"io"
)

type Processor struct {
	Conn net.Conn
}

// 编写一个ServerProcessMes	 根据客户端发送的消息种类不同   决定调用哪个函数
func (this *Processor) serverProcessMes (mes *message.Message) (err error){

	fmt.Println("mes = ", mes)

	switch mes.Type {
		case message.LoginMesType :
			// 处理登录逻辑
			//  创建userProcess 实例
			up := &processes.UserProcess{
				Conn : this.Conn,
			}
			err = up.ServerProcessLogin(mes)
		case message.RegisterMesType :
			// 处理注册逻辑
			up := &processes.UserProcess{
				Conn : this.Conn,
			}
			err = up.ServerProcessRegister(mes)
		case message.SmsMesType :
			smsp := &processes.SmsProcess{}
			smsp.SendGroupMes(mes)
		default :
			fmt.Println("消息类型不存在，无法处理...")

	}
	return 
}


func (this *Processor)  process2() (err error) {

	for {
		// 这里我们将读取的消息，封装成一个函数	readPkg(conn) ()
		// 创建transfer
		tf := &utils.Transfer{
			Conn : this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，我也退出")
				return err
			} else {
				fmt.Println("readPkg err = ", err)
				return err
			}
		}
		//fmt.Println(mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}	
}