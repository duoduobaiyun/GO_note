package main

import "fmt"

func main() {
	//作业二：在作业一的基础上，创建一个函数，该函数接收职员类型，
	//在函数中判断职员的薪水是否大于5000，薪水小于等于5000元，交税数额为0； 如果大于5000，
	//则计算员工要交的税额，并赋值到相应的字段。最后在main函数中调用函数，并接受返回值，打印员工要交的税额。

	a:=information1{
		EmployeeNumber: 05,
		name:           "飞虎",
		section:        "情报队",
		Salary:         20000,
		TaxAmount:      1000,
	}
	fmt.Println(a.EmployeeNumber,a.name,a.section,a.Salary,a.TaxAmount)

  fmt.Println(fun1(a.Salary))

}

func fun1(a int)int {
     b:=a-(a - 5000 )/5
     return b
}

type information1 struct {
	EmployeeNumber  int
	name string
	section string
	Salary  int
	TaxAmount int
}
