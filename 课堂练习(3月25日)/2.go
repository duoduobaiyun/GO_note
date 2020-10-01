package main

import (
	"fmt"
)

func main() {
	/*
	函数的返回值:
	     一个函数的执行结果,返回给函数的调用处,执行结果就叫做函数的返回值。
	return语句:
	     一个函数的定义上有返回值,那么函数中必须使用return语句,将结果返回给调用处
	 */
	//1.设计一个函数,用于求1-10的和,将结果在函数中打印输出
	res:=getsum()
	fmt.Println("1-10的和",res)

	fmt.Println(getsum2())//5050
	fmt.Println("——————————————————————————————")

    res1,res2:=rectangle(5,3)
	fmt.Println("周长:",res1,"面积",res2)
    res3,res4:=rectangle2(5,3)
    fmt.Println("周长",res3,"面积:",res4)

    res5,_:=rectangle2(5,6)
    fmt.Println(res5)
}
//定义一个函数,带回返回值
func getsum()int  {
	sum:=0
	for i:=0;i<=10 ;i++  {
		sum+=1
	}
	return sum
}
func getsum2()(sum int)  {//定义函数时,指明要返回的数据是哪一个
	for i:=1;i<=100 ; i++ {
		sum+=i
	}
	return
}

func rectangle(len,wid float64)(float64,float64)  {
	perimeter:=(len+wid)*2
	area:=len*wid
	return perimeter,area
}

func rectangle2(len,wid float64)(peri  float64,area  float64)  {
	peri = (len + wid)*2
	area = len*wid
	return
}