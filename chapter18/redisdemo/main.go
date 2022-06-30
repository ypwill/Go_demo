package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main()  {


	//通过go 连接redis 操作
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis连接出现错误 ", err)
		return 
	}

	defer conn.Close()

	_, err = conn.Do("set", "name", "北京")
	if err != nil {
		fmt.Println("出现错误 ", err)
		return 
	}
	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("出现错误 ", err)
		return 
	}
	// 因为取出的是r 接口类型 需要转成string
	//nameString := redis.String(r)
	fmt.Println(r)

	fmt.Println("*****************************")

	fmt.Println("操作成功")

}