package main

import "fmt"

func main(){
	var age   int
	fmt.Println(age)

	var name  string
	fmt.Println("name的值是:",name)

	var grade  float32
	fmt.Println(grade)//float 默认值也是0

	var st1,st2  = "刘", "坤"
	fmt.Println(st1,st2)

	var (
		name1 = "刘坤"
		sex = "男"
		adress = "江西省上饶市潘阳县"
	)
	fmt.Println(name1,sex,adress)

	//舍弃变量
	num11,_:=1,2//_符号表示的将该位置上的变量进行舍弃
	//fmt.println(num1,num2)
	fmt.Println(num11)
}
