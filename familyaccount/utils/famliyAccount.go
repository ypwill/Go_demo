package utils

import "fmt"

type FamilyAccount struct{
		// 声明用于接收的key
		key string

		// 声明一个变量退出for循环
		loop bool
	
		// 定义账户余额
		balance float64
		money float64
		note string
		// 收支详情
		// 当有收支时，不停的拼接details
		details string
		origin string
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount {
		key : "",
		loop : true,
		balance : 10000.0,
		money : 0.0,
		note : "",
		details : "收支\t账户金额\t收支金额\t说 明",
		origin : "收支\t账户金额\t收支金额\t说 明",
	}
}


func (this *FamilyAccount) showDetail() {
	fmt.Println("-----------------当前收支明细记录-----------------")
	// !strings.EqualFold(details, origin)
	if this.details != this.origin {
		fmt.Println(this.details)
	} else {
		fmt.Println("没有收支明细,快起使用把")
	}
}

func (this *FamilyAccount) Income() {
	fmt.Printf("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Printf("本次收入说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
}

func (this *FamilyAccount) Pay() {
	fmt.Printf("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.balance < this.money {
		fmt.Println("对不起,没那么多钱!余额不足~")
		//break
		
	}
	this.balance -= this.money
	fmt.Printf("本次支出说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
}

func (this *FamilyAccount) Exit() {
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
		this.loop = false

	}
}

func (this *FamilyAccount) MainMenu() {
	for this.loop {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Printf("请选择(1-4): ")

		fmt.Scanln(&this.key)

		switch this.key {
			case "1":
				this.showDetail()
			case "2":
				this.Income()
			case "3":
				this.Pay()
			case "4":
				this.Exit()
			default :
				fmt.Println("请输入正确的选项..")
			
		}
		if !this.loop {
			break
		}
	}
}