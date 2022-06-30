package main

import (
	"fmt"
	"go_code/customerManage/service"
	"go_code/customerManage/model"
)

type CustomerView struct {
	key string
	loop bool 
	customerSerive *service.CustomerService
}

// 显示所有客户信息
func (this *CustomerView) List() {
	customers := this.customerSerive.List()
	//fmt.Println(customers)
	fmt.Println("--------------------------- 客户列表---------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for _, v := range customers {
		//fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t",v.Id, v.Name, v.Gender, v.Age, v.Phone, v.Email)
		fmt.Println(v.GetInfo())
	}
	fmt.Printf("\n--------------------------- 客户列表完成---------------------------\n")
}

// 添加用户信息
func (this *CustomerView) AddCustomer() {
	//var customer model.Customer
	fmt.Println("--------------------------- 添加客户 ---------------------------")
	fmt.Printf("姓名:")
	name := ""
	//fmt.Scanln(&customer.Name)
	fmt.Scanln(&name)

	fmt.Printf("性别:")
	gender := ""
	//fmt.Scanln(&customer.Gender)
	fmt.Scanln(&gender)
	fmt.Printf("年龄:")
	age := 0
	//fmt.Scanln(&customer.Age)
	fmt.Scanln(&age)
	fmt.Printf("电话:")
	phone := ""
	//fmt.Scanln(&customer.Phone)
	fmt.Scanln(&phone)
	fmt.Printf("邮箱:")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer02(name, gender, age, phone, email)

	this.customerSerive.AddCustomer(customer)
	fmt.Println("--------------------------- 添加完成 ---------------------------")

}


func (this *CustomerView) DeleteCustomer () {
	fmt.Println("--------------------------- 删除客户 ---------------------------")
	var targetId int
	fmt.Printf("请选择待删除客户编号(-1退出):")
	fmt.Scanln(&targetId)
	if targetId == -1 {
		return
	}
	fmt.Printf("确认是否删除(Y/N):")
	var flag string
	fmt.Scanln(&flag)
	if flag == "Y"{
		if this.customerSerive.DeleteCustomer02(targetId) {
			fmt.Printf("\n--------------------------- 删除完成 ---------------------------\n")
		} else {
			fmt.Println("序号不存在，删除失败")
		}
		//this.customerSerive.DeleteCustomer(targetId)
		//fmt.Printf("\n--------------------------- 删除完成 ---------------------------\n")
	}
}

func (this *CustomerView) exit() bool {
	fmt.Printf("确认是否退出(Y/N):")
	var flag string
	fmt.Scanln(&flag)
	if flag == "Y" {
		return false
	} else {
		return true
	}
}


func (this *CustomerView) UpdateCustomer() {
	fmt.Printf("请选择待修改客户编号(-1退出):")
	var index int
	fmt.Scanln(&index)
	flag, oldCustomer := this.customerSerive.GetOldCustomer(index)
	if !flag {
		fmt.Println("客户不存在, 请重新输入")
		return 
	} else {
		var newCustomer model.Customer
		newCustomer.Id = index
		//fmt.Println(oldCustomer)
		fmt.Printf("姓名(%v):", oldCustomer.Name)
		
		fmt.Scanln(&newCustomer.Name)
		if newCustomer.Name == "" {
			newCustomer.Name = oldCustomer.Name
		}

		fmt.Printf("性别(%v):", oldCustomer.Gender)
		fmt.Scanln(&newCustomer.Gender)
		if newCustomer.Gender == "" {
			newCustomer.Gender = oldCustomer.Gender
		}

		fmt.Printf("年龄(%v):", oldCustomer.Age)
		fmt.Scanln(&newCustomer.Age)
		if newCustomer.Age == 0 {
			newCustomer.Age = oldCustomer.Age
		}

		fmt.Printf("电话(%v):", oldCustomer.Phone)
		fmt.Scanln(&newCustomer.Phone)
		if newCustomer.Phone == "" {
			newCustomer.Phone = oldCustomer.Phone
		}

		fmt.Printf("邮箱(%v):", oldCustomer.Email)
		fmt.Scanln(&newCustomer.Email)
		if newCustomer.Email == "" {
			newCustomer.Email = oldCustomer.Email
		}
		//fmt.Println(newCustomer)
		//fmt.Println(index)
		indexPos := this.customerSerive.FindById(index)
		this.customerSerive.UpdateCustomer(indexPos, newCustomer)
	}
}



func (this *CustomerView) MainMenu() {
	for {
		fmt.Println("\n-----------------客户信息管理软件-----------------")
		fmt.Println("                 1 添 加 客 户")
		fmt.Println("                 2 修 改 客 户")
		fmt.Println("                 3 删 除 客 户")
		fmt.Println("                 4 客 户 列 表")
		fmt.Println("                 5 退       出")
		fmt.Printf("请选择(1-5):")
		fmt.Scanln(&this.key)
		switch this.key {
			case "1" :
				this.AddCustomer()
			case "2" :
				this.UpdateCustomer()
			case "3" :
				this.DeleteCustomer()
			case "4" :
				this.List()
			case "5" :
				this.loop = this.exit()
			default :
				fmt.Println("您的输入有误,请重新输入..")
		}
		if !this.loop {
			break;
		}
	}
	fmt.Println("你退出了客户关系管理系统...")
}

func main() {
	customerView := CustomerView{
		key : "",
		loop : true,
	}
	customerView.customerSerive = service.NewCustmoerService()
	customerView.MainMenu()

}