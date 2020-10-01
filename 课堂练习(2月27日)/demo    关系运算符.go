package main

import "fmt"

func main() {
	//关系: 18岁  <  19岁
	//关系运算符:  >  <  >=  <=
	//关系运算符: 运算符运算的结果是true/false,只能是这两个结果当中的一个，go语言中的bool类型
	age1 :=  18
	age2 :=  19
	//英文单词:result,结果
	result  :=age1  >  age2//18  >   19
	fmt.Println(result)
	fmt.Printf("result的类型:%T\n",result)

	result1  := age1 <=  age2
	fmt.Println(result1)

	num1  := 100
	num2  := 100
	result2  := num1  >=  num2
	fmt.Println(result2)
	result3  :=num1 <=  num2
	fmt.Println(result3)

	result4 := num1  == num2
	fmt.Println(result4)

	result5 := num1  != num2
	fmt.Println(result5)
}
