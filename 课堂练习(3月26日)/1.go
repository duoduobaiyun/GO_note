package main

import "fmt"

//全局变量的定义
//num3:=1000//不支持简短定义的写法
var num3 = 1000
func main() {
	/*
	作用域:变量可以使用的范围
	   局部变量:函数内部定义的变量,就叫做局部变量。
	                 变量在哪里定义,就只能在哪个范围使用,超出这个范围,我们认为变量就被销毁了

	   全局变量:函数外部定义的变量,就叫做全局变量。
	 */
	//定义在main函数中,所以n的作用域就是main函数的范围内
	n:=10
	fmt.Println(n)

	if a:=1;a<=10 {
		fmt.Println(a)//1
		fmt.Println(n)//10
	}
	//fmt.Println(a)//不能访问a,出了作用域
	fmt.Println(n)

	if b:=1;b<=10 {
		n:=20
		fmt.Println(b)//1
		fmt.Println(n)//20  根据就近原则选最近的  不同的作用域内,变量名是可以相同的
	}

	fun1()

	fun2()
    fmt.Println("main中访问全局变量",num3)//？
}
func fun1()  {
	//fmt.Println(n)
	num1:=100
	fmt.Println("fun1()函数中:num1",num1)//随着函数的结束,变量都会被销毁的
	num3:=2000
	fmt.Println("fun1()函数,访问全局变量",num3)
}
func fun2()  {
	num1:=200
	fmt.Println(num1)
	fmt.Println("fun2()函数,访问全局变量",num3)
}