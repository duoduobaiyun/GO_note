package main

import (
	"fmt"
)

func main() {
	//1、数据类型转换，在需要得时候，将数据类型进行强制转换
	//数据类型转换得语法:type(value)
	var num    int8//num 在转换成为计算机认识的2进制数的时候，会用8个二进制位来表示
	num = 100

	var num1   int16//num2  在转换成为计算机认识的2进制数的时候，会用16个二进制数来表示
	num1 = 20


	//sum := num +num1
	//sum :=int16(num) + num1
	sum := num + int8(num1)
	fmt.Println(sum)
	//注意:数据类型的转换只能在同类型当中进行转换、重点在数值型
	//age := 18//数值型里面的整形
	//result :=int8(age)
	//fmt.Printf(result)

	grade := 86.4//float
	result :=int8(grade)
	fmt.Printf("数据类型是:%T\n",grade)
	fmt.Printf("数据类型是:%T\n",result)

	var a uint8  =  3//大名  :  张大伟
	var b  byte =  5// 小名  :  二狗子
	sum1  :=  a + b
	fmt.Println(sum1)

	//关于float的小数点位数
	var pai = 3.1415926//
	fmt.Printf("pai的数值是%f\n",pai)
	//保留2位小数
	fmt.Printf("保留两位小数的数值是:%.2f\n",pai)
     fmt.Printf("保留3位小数数值是:%.3f\n",pai)
	fmt.Printf("保留一位小数数值是:%.1f",pai)


	name := "liukun"
	//看一些这个字符串的长度，也就是说字符串当中一共有多少个字符？
	fmt.Println("字符串name的长度是:",len(name))

	adress := "江西省上饶市潘阳县"//一共有11个汉字
	//汉字所占的空间长度跟字母占的空间长度不一样
	length := len(adress)
	fmt.Println("地址啊adress的长度是:",length)

	//截取字符串
	name1 := "tiechuimeimei"//铁锤妹妹
	//如何得到铁锤部分字符串
	tiechui :=name1[0:7]
	fmt.Println(tiechui)

	//截取tiechuimeimei字符串中的meimei
	xiaojiejie :=name1[7:len(name1)]//从m开始切，切斜最后，m的下标是7
	//字符串的长度是:len(name1)
	//初始下标:len(name1)-1
	//结束的下标:len(name1)-1
	fmt.Println(xiaojiejie)

	//如果是切到字符串的末尾，则可以省略不写
	xiaomeimei :=name1[7:]//省略的写法
	fmt.Println(xiaomeimei)

	//如果是从头开始切，也可以不写
	tie :=name1[:7]
	fmt.Println(tie)

	chuimei := name1[3:10]
	fmt.Println(chuimei)
}
