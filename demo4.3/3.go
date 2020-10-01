package main

import "fmt"

func main() {
	//作业三：定义一个函数，实现两数相加，要求最后返回一个指针类型，并在main中调用和接收，最后打印计算后的结果值。
   d:=fun1(2,5)
   fmt.Println(*d)
}


func fun1(a,b  int)*int{
	c:=a+b
	p:=&c
	return p
}
