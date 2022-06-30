package main

import (
	_"fmt"
	"testing"  // 引入testing测试框架	 
)

func TestAddUpper(t *testing.T) {
	  
	res := addUpper(10)
	if res != 55 {
		//fmr.Println("执行出错")
		t.Fatalf("执行出错 %d", res)
	}

	// 如果正确，输出日志
	t.Logf("测试代码执行正确")
}