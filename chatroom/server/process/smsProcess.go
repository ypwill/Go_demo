package processes

import (
	"fmt"
	"net"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
	"encoding/json"
)

type SmsProcess struct {

}

func (this *SmsProcess) SendGroupMes(mes *message.Message) {

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
		if err != nil {
			fmt.Println("json.Unmarshal err = ", err)
			return 
		}
	

	data, err := json.Marshal(mes) 
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}
	
	// 遍历客户端的onlineUsers
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}


func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn : conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败")
	}
	return 
}	