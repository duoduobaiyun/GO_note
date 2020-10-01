package main

import "fmt"

func main() {
	/*
	函数:function
	一、概念:
	    具有特定功能的代码,可以多次调用执行
	二、意义:
	    1、可以避免重复的代码
	    草稿2、增强程序的扩展性
	三、使用:
	    step1:函数的定义,也叫声明
	    step2:函数的调用,就是执行函数中代码的过程
	四、语法
	 */
	//求1 - 10的和
	getsun()
	fmt.Println("hello world...")

	//求1 - 10的和
	getsun()
	fmt.Println("aaa...")

	//求1 - 10的和
	getsun()
}

//定义一个函数:用来求1-10的和
func getsun()  {
	sun:=0
	for i:=1;i<=10 ; i++ {
		sun+=i
	}
	fmt.Printf("1-10的和是:%d\n",sun)
}
