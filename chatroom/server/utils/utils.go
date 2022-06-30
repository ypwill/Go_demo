package utils

import (
	"fmt"
	"net"
	"go_code/chatroom/common/message"
	"encoding/json"
	"encoding/binary"
)

// 这里将这些方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf [8096]byte  // 缓冲
}


func (this *Transfer) ReadPkg()(mes message.Message, err error) {

	fmt.Println("读取客户端发送的数据...")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//fmt.Println("conn.Read err = ", err)
		//err = errors.New("read pkg header error")
		return 
	}
	
	// 根据buf[:4] 转换为一个 unit32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	// conn.Read 函数的含义是  从conn里读 放到buf里面
	n, err := this.Conn.Read(this.Buf[:pkgLen])

	if n != int(pkgLen) || err != nil {
		//fmt.Println("conn.Read err = ", err)
		//err = errors.New("read pkg body error")
		return 
	}

	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return 
	}
	return 
}


func (this *Transfer) WritePkg(data []byte) (err error) {

	// 先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	// 发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err = ", err)
		return 
	}

	// 发送消息
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write err = ", err)
		return 
	}
	return 
}