package main

import (
	"fmt"
)

func main() { //new 内置函数声明结构体。
	emp1 := new(Emp)
	fmt.Printf("emp1: %T，%v , %p\n", emp1, emp1, emp1)

	(*emp1).name = "David"
	(*emp1).age = 30
	(*emp1).sex = 1
	fmt.Println(&emp1)
	// 语法简写形式，语法糖
	emp1.name = "Steven"
	emp1.age = 35
	emp1.sex = 1

	fmt.Println(&emp1)
	fmt.Println("----------------------")
}

// 定义⼀一个结构体
type Emp struct {
	name string
	age  int8
	sex  byte
}
