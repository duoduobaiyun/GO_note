package main

import (
	"fmt"
)

func main()  {
	/*
	变量:variable
	概念:一小块内存，用于存储数据，在程序运行过程中数据可以改变

	使用:
	step1:变量的声明,也叫定义
	step2:变量的访问,复制和运算

	go的特性:
	静态语言:强类型语言
	  go, java, c++, c#..
	动态语言:弱类型语言
	JavaScript, php, python, ruby。。
	 */
	//第一种:定义变量,然后进行赋值
	var  num1  int
	num1 = 30
	fmt.Printf("num1的数值是:%d\n",num1)
	//写在一行
	var  num2  int = 15
	fmt.Printf("num2的数值是:%d\n",num2)

	//第二种:类型推断
	var name  = "王二狗"
	fmt.Printf("类型是:%T,数值是:%s\n",name,name)

	//第三种，简短定义，也叫简短声明
	sun :=100
	fmt.Println(sun)

	//多个变量同时定义
	var a , b , c  int
	a = 1
	b = 2
	c = 3
	fmt.Println(a,b,c)

	var m,n  int =100,200
	fmt.Println(m,n)

	var n1,f1,s1  = 100,3.14,"GO"
	fmt.Println(n1,f1,s1)

	var (
		studentName = "李小花"
		age = 18
		sex = "女"
	)
	 fmt.Printf("学生姓名:%s\n        ,年龄:%d\n        ,性别:%s\n",studentName,age,sex)
}
