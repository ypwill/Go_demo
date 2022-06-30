package model

type student struct {
	Name string
	age int
}

// 因为student首字母小写 因此只能在model包中使用
// 通过使用工厂模式解决该问题

func NewStudent(n string, a int) *student {
	return &student{
		Name : n,
		age : a,
	}
}

// 如果 age 小写 ，则其他包不能直接使用该方法
func (s student) GetAge() int {
	return s.age
}