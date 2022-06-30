package main

import (
	"fmt"
	"go_code/chatroom/server/model"
	"net"
	"time"
)

//   {\"userid\":100,\"userPwd\":\"123456\",\"userName\":\"tom\"}"

// func readPkg(conn net.Conn)(mes message.Message, err error) {

// 	buf := make([]byte, 8096)
// 	fmt.Println("读取客户端发送的数据...")
// 	_, err = conn.Read(buf[:4])
// 	if err != nil {
// 		//fmt.Println("conn.Read err = ", err)
// 		//err = errors.New("read pkg header error")
// 		return
// 	}

// 	// 根据buf[:4] 转换为一个 unit32
// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(buf[:4])

// 	// conn.Read 函数的含义是  从conn里读 放到buf里面
// 	n, err := conn.Read(buf[:pkgLen])

// 	if n != int(pkgLen) || err != nil {
// 		//fmt.Println("conn.Read err = ", err)
// 		//err = errors.New("read pkg body error")
// 		return
// 	}

// 	err = json.Unmarshal(buf[:pkgLen], &mes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal err = ", err)
// 		return
// 	}

// 	return
// }

// 编写一个ServerProcessMes	 根据客户端发送的消息种类不同   决定调用哪个函数
// func serverProcessMes(conn net.Conn, mes *message.Message) (err error){

// 	switch mes.Type {
// 		case message.LoginMesType :
// 			// 处理登录逻辑
// 			err = serverProcessLogin(conn, mes)
// 		case message.RegisterMesType :
// 			// 处理注册逻辑
// 		default :
// 			fmt.Println("消息类型不存在，无法处理...")

// 	}
// 	return
// }

// 编写一个 ServerProcessLogin 函数  专门处理登录逻辑
// func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
// 	var loginMes message.LoginMes

// 	err = json.Unmarshal([]byte(mes.Data), &loginMes)
// 	if err!= nil {
// 		fmt.Println("json.Unmarshal fail err = ", err)
// 		return
// 	}

// 	// 声明 resMes    构建message  结构体并返回
// 	var resMes message.Message
// 	resMes.Type = message.LoginResMesType

// 	var loginResMes message.LoginResMes

// 	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
// 		// 合法
// 		loginResMes.Code = 200
// 	} else {
// 		// 不合法
// 		loginResMes.Code = 500
// 		loginResMes.Error = "该用户不存在，请注册后再使用..."
// 	}

// 	data, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal fail err = ", err)
// 		return
// 	}
// 	// 将序列化后的data 赋值给 构建的 message   然后再序列化
// 	resMes.Data = string(data)

// 	data, err = json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal fail err = ", err)
// 		return
// 	}
// 	// 发送 data  我们将其封装到 writePkg函数
// 	err = writePkg(conn, data)
// 	return
// }

// func writePkg(conn net.Conn, data []byte) (err error) {

// 	// 先发送一个长度给对方
// 	var pkgLen uint32
// 	pkgLen = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
// 	// 发送长度
// 	n, err := conn.Write(buf[:4])
// 	if n != 4 || err != nil {
// 		fmt.Println("conn.Write err = ", err)
// 		return
// 	}

// 	// 发送消息
// 	n, err = conn.Write(data)
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("conn.Write err = ", err)
// 		return
// 	}

// 	return

// }

func process(conn net.Conn) {
	defer conn.Close()
	// 循环接收客户端发送的消息

	// for {

	// 	// 这里我们将读取的消息，封装成一个函数	readPkg(conn) ()
	// 	mes, err := readPkg(conn)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			fmt.Println("客户端退出，我也退出")
	// 			return
	// 		} else {
	// 			fmt.Println("readPkg err = ", err)
	// 			return
	// 		}
	// 	}
	// 	//fmt.Println(mes)
	// 	err = serverProcessMes(conn, &mes)
	// 	if err != nil {
	// 		return
	// 	}
	// }
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("processor.process2 fail err = ", err)
		return
	}
}

// 这里我们编写一个函数，完成对UserDao的初始化任务
func initUserDao() {
	// 这里的 pool 本身就是全局变量  在redis.go 并且已经初始化
	model.MyUserDao = model.NewUserDao(pool)
}

func init() {
	initPool("127.0.0.1:6379", 16, 0, 300*time.Second)

	initUserDao()
}

func main() {

	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
	}

	//一旦监听成功 就等待客户端来连接
	for {
		fmt.Println("等待客户端来连接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err = ", err)
		}

		go process(conn)
	}
}
