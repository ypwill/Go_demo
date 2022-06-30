package main

import "fmt"

// 声明一个接口
type Usb interface {
	// 声明两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {

}

// 让Phone实现 Usb方法
func (p Phone) Start(){
	fmt.Println("手机开始充电")
}

func (p Phone) Stop(){
	fmt.Println("手机停止充电")
}

type Camera struct {

}

// 让Camera实现 Usb方法
func (c Camera) Start(){
	fmt.Println("相机开始充电")
}

func (c Camera) Stop(){
	fmt.Println("相机停止充电")
}

type Computer struct {

}

// 编写一个working方法 接收一个usb接口类型
// 
func (c Computer) working(usb Usb) {
	usb.Start()
	usb.Stop() 
}



func main()  {

	com := Computer{}
	phone := Phone{}
	camera := Camera{}

	// 关键点
	com.working(phone)
	com.working(camera)
	
}