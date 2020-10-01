package main

import "fmt"

func main() {

fmt.Println(fun1([6]int{1,2,3,4,5,6}))
fmt.Println(fun1([6]int{5,3,4,2,7,8}))

s:=[]int{5,2,3,6,1}
reverse2(s)
fmt.Println(s)

c:=&[6]int{2,5,3,7,1,6}
fun2(c)
fmt.Println(fun2(c))
fmt.Printf("%T",fun2(c))

}

func fun1(a [6]int) [6]int  {
	//作业一：编写一个函数，该函数的作用是将数组进行翻转，并返回数组。附程源码和效果运行截图。
	for i:=0;i<len(a)/2 ;i++  {
		a[i],a[len(a) - 1 - i]=a[len(a) - 1 - i],a[i]
	}
	return a
	//fmt.Println("——————————————————————————————————————————————————")
}

func reverse2(sl []int){
	for index := 0; index < len(sl) /2; index++{//从第一个元素开始，遍历数组的前半部分
		//a = arr[index]
		//arr[index] = arr[len(arr) - 1 - index]
		//arr[len(arr) - 1 - index] = a
		sl[index],sl[len(sl) - 1 - index] = sl[len(sl) - 1 - index],sl[index]
	}
}

func fun2(b *[6]int) *[6]int  {
	//作业二：编写一个函数，该函数接收指针类型的数组，在函数中实现对数组的从小到大的排序，
	//最后返回一个指针数组。附程序源码和运行效果。
	for i:=0;i<len(b) - 1 ; i++ {
		for j:=i+1;j<len(b) ; j++ {
			if b[i] > b[j]  {
               b[i],b[j]=b[j],b[i]
			}
		}
	}
	return b
}

