package model

import (
	_"fmt"
	"net"
	"go_code/chatroom/common/message"
)

type CurUser struct {
	Conn net.Conn
	message.User
}

