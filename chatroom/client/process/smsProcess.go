package process

import (
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/client/utils"
	"encoding/json"
)

type SmsProcess struct {

}

func (this *SmsProcess) SendGroupMes(content string) (err error){

	var mes message.Message
	mes.Type = message.SmsMesType

	// 创建 一个 SmsMes
	var smsMes message.SmsMes
	smsMes.Content = content // 内容
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
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
		Conn : CurUser.Conn,
	}
	// 发送data给服务器端
	err = tf.WritePkg(data) 
	if err != nil {
		fmt.Println("注册发送信息错误 err = ", err)
		return 
	}
	return


}