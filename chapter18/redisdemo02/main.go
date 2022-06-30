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

	_, err = conn.Do("hset", "user01", "name", " 小明")
	if err != nil {
		fmt.Println("出现错误 ", err)
		return 
	}
	_, err = conn.Do("hset", "user01", "age", 11)
	if err != nil {
		fmt.Println("出现错误 ", err)
		return 
	}
	r, err := redis.String(conn.Do("hget", "user01", "age"))
	if err != nil {
		fmt.Println("出现错误 ", err)
		return 
	}
	fmt.Println(r)
	
	fmt.Println("操作成功")

}