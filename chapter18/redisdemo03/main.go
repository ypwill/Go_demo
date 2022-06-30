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

	_, err = conn.Do("mset", "name", "xiaoming", "age", "11")
	if err != nil {
		fmt.Println("出现错误 ", err)
		return 
	}
	r, err := redis.Strings(conn.Do("mget", "name", "age"))
	if err != nil {
		fmt.Println("出现错误 ", err)
		return 
	}
	
	for _, v := range r {
		fmt.Println(v)
	}
	
	
	fmt.Println("操作成功")

}