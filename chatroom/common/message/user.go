package message

type User struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
	UserStatus int `json:"userStatus"`  // 用户状态... 在线、离线 等
	Sex string `json:"sex"`
}