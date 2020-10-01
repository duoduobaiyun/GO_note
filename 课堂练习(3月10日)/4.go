package main

import "fmt"

func main() {
	//语法:
	//var 数组名 [长度] 数据类型
	//var 数组名 = [长度] 数据类型{元素1，元素2.。。}
	//数组名   :=[...] 数据类型 {元素...}




	var num1  int
	num1 = 100
	num1 = 200
	fmt.Println(num1)

	//创建数组
	var arr1 [4] int
	//数组的访问
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	fmt.Println(arr1[0])
	fmt.Println(arr1[2])

	fmt.Println("数组的长度:",len(arr1))//容器中实际存储的数据量
	fmt.Println("数组的容量",cap(arr1))//容器中能够存储的最大的数量
	//因为数组定长，长度和容量相同
	arr1[0] = 100
	fmt.Println(arr1[0])

	//数组的其他创建方式
	var a [4] int//同 var a = [4] int
	fmt.Println(a)

	var b  = [4]int{1,2,3,4}
	fmt.Println(b)

	var c  = [5]int{1,2,3}
	fmt.Println(c)

	var d  = [5]int{1:1,3:2}//左边的是下标,右边的是赋予左边值的数字
	fmt.Println(d)

	var e = [5]string{"rose","王二狗","ruby"}//它的零值是空的字符串
	fmt.Println(e)

	f :=[...]int{1,2,3,4,5}
	fmt.Println(f)
	fmt.Println(len(f))

	g :=[...]int{1:3,6:5}
	fmt.Println(g)
	fmt.Println(len(g))
}
