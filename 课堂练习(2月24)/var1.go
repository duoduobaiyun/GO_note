package main

import "fmt"

func  main(){
	//var age int
	//age = 18
	//fmt.Println(age)

	var a,b,c int //3,4,5
	//3、4、5
	a = 3
	b = 4
	c = 5
	fmt.Println(a,b,c)

	a = 6
	b = 5
	c = 10
	fmt.Println(a,b,c)

	//a = "liu"//表示 语法错误
	// a = 刘坤   双引号引起来的表示是字符串

	var st1 ,st2  string = "刘","坤"
	fmt.Println(st1,st2)

	var  name ,age,adress  = "刘坤", "18","江西省鄱阳县"
	fmt.Println(name,age,adress)
	fmt.Printf("name的数据类型:%T\n",name)

	fmt.Printf("name的数据类型是:%s,其数据类型是:%T\n",name,name)
	fmt.Printf("age的数值值是:%d,其数据类型是:%T\n",age,age)

	num1,num2:=1,3
	fmt.Println(num1 + num2)
	//注意1:只有已近被声明才能被使用
	//fmt.Println(num3)//没有声明和定义的变量不能够被使用

	num1,num2 = 5,8//多个变量的赋值
	fmt.Println(num1,num2)

	num1,num3 :=10,20
	fmt.Println(num1,num3)

	num1 :=50//为什么报错？ 原因:=表示的是新定义
	fmt.Println(num1)
}