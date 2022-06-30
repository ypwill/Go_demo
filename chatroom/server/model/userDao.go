package model

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"go_code/chatroom/common/message"
)

// 我们在服务器启动后，就初始化一个userDao实例
// 把他做为全局的变量，在需要和redis操作时，就直接使用即可

var (
	MyUserDao *UserDao
)


// 定义一个userDao结构体 完成对user的各种操作
type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式 创建一个UserDao 实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {

	userDao = &UserDao{
		pool : pool,
	}
	return 

} 


func (this *UserDao) getUserById (conn redis.Conn, id int)(user *User, err error) {

	res, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return 
	}

	user = &User{}
	// 这里 我们放入 redis的是 序列化后的data 我们先需要反序列化
	err = json.Unmarshal([]byte(res), &user) 
	if err != nil {
		fmt.Println("json.Unmarshal fail err = ", err)
		return 
	}
	return 
}


// 	Login 完成登陆的校验
// 如果 id 和 pwd 都正确  返回一个user实例    用户的id或者pwd错误 返回错误信息
func (this *UserDao) Login (userId int, userPwd string)(user *User, err error) {

	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return 
	}

	// 这个时候只是证明用户id是对的  还需要校验密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return 
	}

	return 
}


func (this *UserDao) Register (user *message.User)(err error) {

	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return 
	}
	// 到这里 说明用户还没有注册过  则可以完成注册
	data, err := json.Marshal(user) // 序列化    放到redis里面
	if err != nil {
		fmt.Println("json.Marshal fail err = ", err)
	}
	_, err = conn.Do("hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存用户注册 fail err = ", err)
	}



	return 
}