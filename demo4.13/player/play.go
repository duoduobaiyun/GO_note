package player

import "fmt"

type Run struct {
	Name   string
	Sex    string
	SportEvents  string
	grade   int
}

func (a *Run) Setgrade(g int)  {
	a.grade = g
}

//func (a *Run) Getgrade() int  {
//	return a.grade
//}
//可以省略

func (a Run) Run()  {
	fmt.Printf("姓名:%s,性别:%s,体育项目:%s,成绩:%d",a.Name,a.Sex,a.SportEvents,a.grade)
}