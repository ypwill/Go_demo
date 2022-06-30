package main

import (
	"fmt"
	"time"
	"strconv"
)

// 时间日期相关函数
func main() {

	// 1.获取当前时间
	now := time.Now()
	fmt.Println(now)   // 2022-06-16 18:52:37.657479 +0800 CST m=+0.004000701

	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(int(now.Month()))
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// 格式化日期时间
	// 1.使用 printf格式化输出
	// 2.使用Format
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))

	// 每隔一秒打印一个数字 
	// i := 0;
	// for {
	// 	i++;
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second)
	// 	if (i > 10){
	// 		break
	// 	}
	// }

	// 获取当前unix 时间戳
	i := now.Unix()
	fmt.Println(i)


	
	start := time.Now().Unix()
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
	end := time.Now().Unix()
	fmt.Println(end - start)

	



}