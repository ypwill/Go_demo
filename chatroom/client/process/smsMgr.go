package process

import (
	"fmt"
	"encoding/json"
	"go_code/chatroom/common/message"
)

func outputGroupMes(mes *message.Message) {
	// 先反序列化mes.data
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return 
	}
	

	// content := smsMes.Content
	info := fmt.Sprintf("用户id:\t%d对大家说:\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()

}