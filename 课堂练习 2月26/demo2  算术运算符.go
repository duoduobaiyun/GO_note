package main

import (
	"fmt"
)

func main() {
	//数学中的数学运算:加法、减法、乘法、除法
	//运算符的概念:参与变量运算的符号，称之为运算符
	a := 3 //整数3
	b := 5 //整数5
    sum := a + b
    fmt.Println(sum)
    sub := a - b
    fmt.Println(sub)
    mul := a * b
    fmt.Println(mul)
    div :=a / b
    fmt.Println(div)
    //5/3 = 1.66
    //5/3 =1.....草稿2  商1余2
    mod :=b%a
    fmt.Println(mod)

    age := 18
    fmt.Println("铁锤妹妹的年龄",age)
    age++//自动+1 等同于age = age + 1
    fmt.Println("过了年后,铁锤妹妹的年龄:",age)

    num :=2020
    fmt.Println("当前的年份是:",num)
    num--//自动减一， 等同于num = num -1
    fmt.Println("穿越后的年份是:",num)
}
