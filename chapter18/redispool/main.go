package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle : 8,
		MaxActive : 0,
		IdleTimeout : 100,
		Dial : func()(redis.Conn, error){
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main()  {


	//通过go 连接redis 操作
	// conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	// if err != nil {
	// 	fmt.Println("redis连接出现错误 ", err)
	// 	return 
	// }

	// defer conn.Close()

	conn := pool.Get()

	_, err := conn.Do("set", "name", "AAA")
	if err != nil {
		fmt.Println(err)
	}

	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(r)
	defer pool.Close()
	

}