package main

import "fmt"

func main() {
	a := 8
	b := 10
	c := 12

	//1、关系运算符的结果是bool:true、false
	//草稿2、使用&&和||对整个表达式进行分割 a < b 、 b/草稿2 >3、a * 草稿2 >  b、c != 12
	//3、分别计算结果:true、true、true、false
	result1 := a < b && b/2 >3 && a*2 >b || c != 12
	fmt.Println("运算结果是:",result1)

	result2 := a < b && b/2 > 3 && a * 2 < b || c !=12
	fmt.Println("result2的运算结果是:",result2)
}
