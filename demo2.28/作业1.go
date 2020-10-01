package main

import "fmt"

func main() {

	num1 := 2
	num2 := 7

	res1 := num1 & num2 // 按位与的操作
	fmt.Println(res1)

	//按位或的操作
	age1 := 4
	age2 := 7
	res2 := age1 | age2
	fmt.Println(res2)

	//按位异或
	a := 6
	b := 9
	res3 := a ^ b
	fmt.Println(res3) // 15


	num1 = 3
	num2 = 5
	res4 := num1 ^ num2
	fmt.Println(res4) // 6


	num1 = 3 // 11
	res5 := num1 << 2 // 左移两位
	fmt.Println(res5) // 12

	num1 = 7 //
	res6 := num1 >> 2 // 向右移动2位
	fmt.Println(res6)

}
