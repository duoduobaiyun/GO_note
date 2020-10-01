package main

import "fmt"

func main() {
	a := 3
	b := 4


	//第一小题
	res1 := a < b && b / 2 == 1 && a % 3 != 0
	fmt.Println("res1的值是:",res1)


	res2 := (a+b)*3 < a<<2 || (a-b) >0
	fmt.Println("res2的值是:",res2)

	d,c := "hello","world"
	fmt.Println(d,c)
}
