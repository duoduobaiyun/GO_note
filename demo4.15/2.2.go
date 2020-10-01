package main

import "fmt"

func main() {
	//作业二：定义长度为6的数组并进行赋初始值。
	//使用for循环访问0到10的数组下标上的元素。
	//处理程序可能遇到的异常，处理方式主动终止程序执行，并给出异常的原因。
	a:=[6]int{1,2,3,4,5,6}
	for i:= 0;i<10 ; i++ {
		if i>=6 {
			fmt.Println(a[i])
			panic("数组越界")
		}
	}
}
