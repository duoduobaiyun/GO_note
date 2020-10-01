package main

import "fmt"

func main() {
	/*
	可变参数://可变参数就是一个切片,切片就是一个长度可变的数组
	   概念:一个函数的参数的类型确定,但是个数不确定,就可以使用可变参数
	   语法:
	        参数名...参数的类型

	        对于函数,可变量参数相当于一个切片
	        调用函数的时候,可以传入0-多个参数

	        println(),printf(),print()//接受可变参,作为参数的
	        append()//也是可变参数
	    注意事项:
	A:如果一个函数的参数是可变参数,同时还有其他的参数,可变参数要放在参数列表的最后
	B:一个函数的参数列表中最多只能有一个可变参数
	 */
	//1.求图
	getSum()
	getSum(1,2,3,4,5)
	getSum(1,2,3,4,5,6,7,8,9,10)
	//草稿2.切片
	s1:=[]int{1,2,3,4,5}
	getSum(s1...)
}
func getSum(nums...int)  {
	//fmt.Printf("%T\n",nums)//[]int
	sun:=0
		for i:=0;i<len(nums) ; i++ {
             sun += nums[i]
		}
		fmt.Println("总和是:",sun)
	}

func fun1(s1,s2 string,nums...float64)  {
	
}