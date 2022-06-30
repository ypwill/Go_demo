package main

import "fmt"


type Student struct {
	Name string
	Age int
	Score int
}


func (stu *Student) showInfo() {
	fmt.Printf("%v %v %v", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore (score int) {
	stu.Score = score
}

// 小学生
type Pupil struct {
	Student
}

type Graduate struct {
	Student
}


func (p *Pupil) test() {
	fmt.Println("小学生在考试")
} 

func (g *Graduate) test() {
	fmt.Println("大学生在考试")
} 


func main() {

	puile01 := Pupil{}
	puile01.Student.Name = "小明"
	puile01.Student.Age = 18
	puile01.Student.SetScore(11)
	puile01.test()
	puile01.showInfo()

	fmt.Println()

	puile02 := Pupil{}
	puile02.Name = "小二"
	puile02.Age = 18
	puile02.SetScore(22)
	puile02.test()
	puile02.showInfo()

	
}