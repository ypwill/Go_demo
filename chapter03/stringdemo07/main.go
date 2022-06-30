package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
)

func main() {
	
	string := "hello world";
	fmt.Println(string)

	str1 := ` package main

	// import "fmt"
	// import "unsafe"
	import (
		"fmt"
	)
	`
	fmt.Println(str1)

	// 字符串 使用 + 拼接时  要留在上一行（go 默认会在每行末尾加分号，所以得放在上一行）
	str2 := "hello"+ 
	"hello" + "hello" + "hello" + "hello" +
	 "hello" + "hello" + "hello"
	fmt.Println(str2)

}