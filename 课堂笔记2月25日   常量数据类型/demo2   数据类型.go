package main

import (
	"fmt"
)

func main() {
	var num  int8
	fmt.Printf("num的数值:%d,数据类型是%T\n",num,num)

	var age  int16
	age = 20
	fmt.Printf("num的数值是:%d,数据类型是:%T\n",age,age)

	var result  bool
	result = true
	fmt.Println(result)

	result = false
	fmt.Println(result)
}
