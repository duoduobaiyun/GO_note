package main

import "fmt"

func main() {
	/**
	  编写一个函数，该函数接收数组的指针，在函数中实现对数组的从小到大的排序，最后返回一个数组指针

	  指针：存储另一个变量的内存地址
	  格式：*Type ：type可选值包括：int、string、bool、float、数组

	  arr := [3]int{3,5,1}
	  p1 := &arr

	*/
	arr := [3]int{3,5,1}// arr的类型：[3]int ，3是数组类型的一部分
	p1 := &arr//&获取地址  因此: 表示一个数组指针
	//数组指针的意思是：p1本质上是一个指针，只不过这个指针的类型是数组类型
	//打印p1的类型
	fmt.Printf("%T\n",p1)// *[3]int

	num := 100// int
	p2 := &num //
	fmt.Printf("%T\n",p2) // *int, 表示整形指针
	//p2整形指针的意思：p2本质上是一个指针，只不过这个指针的类型是整数类型

	num1 := 1
	num2 := 3
	num3 := 5

	var arr1 [3]*int// 指针数组：本质上是一个数组，数组当中的元素值是一个地址
	arr1[0] = &num1
	arr1[1] = &num2
	arr1[2] = &num3
	fmt.Println(arr1)//[0xc00005e0b0 0xc00005e0b8 0xc00005e0c0] 说明 arr1是一个数组，每一个元素的值是一个地址

	fmt.Println(" =========================== ")
	arr2 := [4]int{1,2,3,4}//数组  [4]int
	fmt.Printf("数组的地址:%p\n",&arr2)
	p3 := &arr2
	fmt.Printf("%p\n",p3)
	fmt.Println(&arr2[0])//打印第0个元素的地址
	//fmt.Print(p3)// &[1 草稿2 3 4]

	var p *[4]int
	p = &arr2
	fmt.Printf("%p\n",p)

	fmt.Println(" ------------------------------")
	//指针数组
	//a := 10
	//b := 12
	//c := 15
	//arr3 := [3]*int{&a,&b,&c}
	//fmt.Println(arr3)

	//
	arr4 := [5]int{12,23,31,9,5}
	sort(&arr4)// &arr4表示arr4的地址
	fmt.Println(sort(&arr4))
}

/**
  编写一个函数，该函数接收数组的指针，在函数中实现对数组的从小到大的排序，最后返回一个数组指针
  1、写一个函数: 新定义一个函数
  草稿2、函数：接收一个数组指针，对数组进行从小打到排序
  3、函数的返回值：返回数组指针
*/
//sort属于 函数指针 还是属于 指针函数？
//1、确定sort是一个函数
//草稿2、sort的返回值类型 *[5]int，即返回值类型是 数组指针
//3、因此sort是指针函数。
func sort(p *[5]int) *[5]int{
	//排序
	//排序：冒泡，选择
	//冒泡：从第一个元素开始比，每个元素和下一个元素比较，如果前面的比后面的要大，则交换位置
	for i := 0; i < len(p) -1;i++{//
		for j := i+1; j < len(p); j++{
			if p[i] > p[j]{
				p[i],p[j] = p[j],p[i]
			}
		}
	}
	return p
}
