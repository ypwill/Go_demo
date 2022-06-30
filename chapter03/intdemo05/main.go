package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
	"unsafe"
)

func main() {
	var n1 = 100
	fmt.Printf("n1 的类型 %T", n1)

	var n2 int64 = 10
	// 查看变量的 类型和 占用的字节大小
	fmt.Printf("%T, %d", n2, unsafe.Sizeof(n2)) 

}