package main

import "fmt"

func main()  {
	//定义一个变量age年龄，表示一个人的年龄，并且赋值为18岁
	//var age int=18
	age:=18//变量声明的第二种写法，省略var和数据类型
	fmt.Println("我的年龄是:",age)
	//想看一下age是什么样的数据类型？
	fmt.Printf("age的数据类型是:%T\n",age)//%T的数据类型
	name:="刘坤"
	fmt.Printf("name的数据类型是:%T",name)//string 代表字符串
}
