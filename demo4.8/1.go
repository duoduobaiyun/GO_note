package main

import "fmt"

func main() {
	//作业一：公司里的职员，可以描述职员的信息包括：员工编号，姓名，部门，薪水，交税数额。
	//要求：编程创建结构体。并创建几个具体的职员信息，交税数额均设置为默认值0，然后打印输出这些职员的信息。

	a:=information{
		EmployeeNumber: 05,
		name:           "飞虎",
		section:        "情报队",
		Salary:         20000,
		TaxAmount:      1000,
	}
	fmt.Println(a.EmployeeNumber,a.name,a.section,a.Salary,a.TaxAmount)
}

type information struct {
	EmployeeNumber  int
	name string
	section string
	Salary  int
	TaxAmount int
}