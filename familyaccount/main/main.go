package main

import (
	"go_code/familyaccount/utils"
	"fmt"
)

func main() {
	fmt.Println("这个是面向对象方式完成")
	familyaccount := utils.NewFamilyAccount()
	familyaccount.MainMenu()
}