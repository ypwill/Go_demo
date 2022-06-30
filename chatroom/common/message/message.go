package message

const (
	LoginMesType			= "LoginMes"
	LoginResMesType			= "LoginResMes"
	RegisterMesType			= "RegisterMes"
	RegisterResMesType		= "RegisterResMes"
	NotifyUserStatusMesType	= "NotifyUserStatusMes"   // 服务器主动发送的推送  
	SmsMesType				= "SmsMes"
)

// 定义用户在线状态的常量
const (
	UserOnline = iota
	Useroffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// 定义两个消息 登录相关  后续再增加

type LoginMes struct {
	UserId int `json:"userid"`
	UserPwd string `json:"userPwd"`
	UserName string	`json:"userName"`
}

type LoginResMes struct {
	Code int `json:"code"` // 500 表示用户未注册  200 表示登录成功
	// 登录返回的消息 带上 在线用户信息
	UsersId []int
	Error string `json:"error"`
}


type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code int `json:"code"` // 400 表示用户已经占用  200 表示登录成功
	Error string `json:"error"`
}

// 为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`	// 用户id
	Status int `json:"status"` // 用户状态
}


type SmsMes struct {
	Content string `json:"content"`
	User
}