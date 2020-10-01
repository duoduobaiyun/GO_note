package main

import "fmt"

func main() {
	//作业二：定义一个长度为6的字符串数组，并进行元素的初始化赋值；
	//要求：将字符串数组中各元素的地址保存到一个指针数组中，并打印该指针数组的元素值。
	//最后通过操作指针数组元素修改原字符串数组中最后一个元素的元素值


	a:=1
	b:=2
	c:=3
	d:=4
	e:=5
	f:=6
	arr1:=[6]int{a,b,c,d,e,f}
	arr2:=[6]*int{&a,&b,&c,&d,&e,&f}
    fmt.Println(arr1)
	fmt.Println(arr2)
    *arr2[5]=100
    fmt.Println(arr1)
    fmt.Println(f)



}
