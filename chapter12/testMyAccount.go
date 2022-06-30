package main

import (
	"fmt"
	_"strings"
)

func main() {

	// 声明用于接收的key
	key := ""

	// 声明一个变量退出for循环
	loop := true

	// 定义账户余额
	balance := 10000.0
	money := 0.0
	note := ""
	// 收支详情
	// 当有收支时，不停的拼接details
	details := "收支\t账户金额\t收支金额\t说 明"
	origin := details


	// 显示主菜单
	for loop {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Printf("请选择(1-4): ")

		fmt.Scanln(&key)

		switch key {
			case "1":
				fmt.Println("-----------------当前收支明细记录-----------------")
				// !strings.EqualFold(details, origin)
				if details != origin {
					fmt.Println(details)
				} else {
					fmt.Println("没有收支明细,快起使用把")
				}
			case "2":
				fmt.Printf("本次收入金额：")
				fmt.Scanln(&money)
				balance += money
				fmt.Printf("本次收入说明：")
				fmt.Scanln(&note)
				details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)
			case "3":
				fmt.Printf("本次支出金额：")
				fmt.Scanln(&money)
				if balance < money {
					fmt.Println("对不起,没那么多钱!余额不足~")
					break
					
				}
				balance -= money
				fmt.Printf("本次支出说明：")
				fmt.Scanln(&note)
				details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)
			case "4":
				fmt.Printf("你确定要退出吗？ y/n:")
				choice := ""
				for {
					fmt.Scanln(&choice)
					if choice == "y" ||  choice == "n" {
						break
					}
					fmt.Printf("你的输入有误, 请重新输入y/n:")
				}
				if choice == "y" {
					loop = false

				}
			default :
				fmt.Println("请输入正确的选项..")
			
		}
		if !loop {
			break
		}
	}

	fmt.Println("你退出了家庭记账软件的使用...")


}