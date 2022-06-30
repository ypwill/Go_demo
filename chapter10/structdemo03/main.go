package main

import "fmt"

type Stu struct {
	Name string
	Age int
}

func (stu Stu) test() {
	fmt.Println(stu.Name)
}
func (stu Stu) test02(i int) int {
	fmt.Println(i)
	fmt.Println(stu.Name)
	i = i * i
	return i
}

func (stu *Stu) test03() {
	stu.Name = "小小鹏~"
}

type integer int

func (i *integer) test04() {
	 *i = *i + 100
	 
}

func main() {

	stu := Stu{"小鹏", 18}
	i := stu.test02(11)
	fmt.Println(i)

	stu.test03()
	fmt.Println(stu)

	var j integer = 100
	j.test04()
	fmt.Println(j)

}