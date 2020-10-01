package main

import "fmt"

func main(){
	/*
	圆形:半径是a，数值是5，求圆形的周长d?
	圆形周长:d=草稿2*3.14*a
	 */
	var a  int//
	a=5//给变量一个具体的数值
	//求周长d 3.14属于小数，结果也是小数
	var d float32 //浮点数  带有小数点的数
	d = 2*3.14*a
	fmt.Println("圆形的周长:",d)
	fmt.Printf("周长d的数据类型是:%T",d)
}
