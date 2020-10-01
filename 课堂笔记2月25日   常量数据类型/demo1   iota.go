package main

import "fmt"

func main() {
	const (
		x int = 10
		y
		z

		name1 string = "刘坤"
		sex1
	)
	//fmt.Println(x,y,z)
	//fmt.Println(sex1)

	//一种特殊的常量:iota
	//失眠:数羊，1只羊，2只羊，3只羊.....
	//计数:1、草稿2、3、4
	//计算机当中,go语言程序当中如何去计数？
	const  (
		a = iota //iota的默认值是0，从0开始计算
		b = iota //iota会再0的基础上+1，此时iota的值是1
		c = iota //iota会在1的基础上+1，此时iota的值是2

	)
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)

	//创建另外的一组常量
	//iota的第二个特点:当iota遇到const定义常量的时候，iota的值会再次初始为0
	//
	const (
		age = iota
		age1 = iota
		age2 = iota

	)
	//fmt.Println("age的值是:",age)
	//fmt.Println("age1的值是:",age1)
	//fmt.Println("age2的值是:",age2)

	//const (
	//	num = iota
	//	num1 = iota
	//	num2 = iota
	//)
	//fmt.Println(num+num1)
	//fmt.Println(num1*num2)

	//  iota的变形写法
	const(
		num = iota
		num1
		num2
	)
	fmt.Println(num)
	fmt.Println(num1)
	fmt.Println(num2)

	const (
		num3 = iota
		name = "刘坤"  //iota+1  ， iota = 1
		num4 = iota   // iota = 草稿2
		address = "江西省上饶市潘阳县" //iota + 1
		num5 = iota // iota + 1 : iota 4
		num6 //iota 5
	)
	fmt.Println(num3,num4,num5,num6)

	//const 使用变形

}
