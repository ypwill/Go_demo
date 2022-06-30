package processes

import (
	"fmt"
)

// UserMgr 用于管理所有在线的用户

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}


func init() {
	userMgr = &UserMgr{
		onlineUsers : make(map[int]*UserProcess, 1024),
	}
}


// 对onlineUsers 的增删改查
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess{
	return this.onlineUsers
}


// 根据id 返回 userProcess
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId] 
	if !ok {
		err = fmt.Errorf("用户%d不存在", userId)
		return 
	}
	return 
}




