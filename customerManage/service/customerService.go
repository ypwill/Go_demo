package service

import (
	"go_code/customerManage/model"
	"fmt"
)

// 创建customerService 完成对 Customer的操作
type CustomerService struct {
	customers []model.Customer
	customerNum int
}

func NewCustmoerService() *CustomerService {
	customerSerive := &CustomerService{}
	customerSerive.customerNum = 1
	customer := model.NewCustomer(1, "小明", "男", 18, "17746611780", "834@qq.com")
	customerSerive.customers = append(customerSerive.customers, customer) 
	return customerSerive
} 

// 返回客户切片
func (this *CustomerService) List() []model.Customer{
	return this.customers
}

func (this *CustomerService) AddCustomer(customer model.Customer) {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
}

func (this *CustomerService) FindById (id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}

// 根据id删除客户
func (this *CustomerService) DeleteCustomer02 (targetId int) bool{
	index := this.FindById(targetId)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}


func (this *CustomerService) DeleteCustomer (targetId int){
	for i, v := range this.customers {
		if v.Id == targetId {
			this.customers = append(this.customers[:i], this.customers[i+1:]...)
		}
	}

}

func (this *CustomerService) GetOldCustomer(id int) (bool, model.Customer){
	index := this.FindById(id)
	var customer model.Customer
	if index == -1 {
		return false, customer
	}
	for i, v := range this.customers {
		if i == index {
			customer = v
		}
	}
	return true, customer
}

func (this *CustomerService) UpdateCustomer(i int, customer model.Customer) {
	for k, _ := range this.customers {
		if k == i {
			fmt.Println(1)
			this.customers[k] = customer
		}
	}
}
 
