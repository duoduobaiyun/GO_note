package main

import "fmt"

func main() {
	var a int
	fmt.Println("请输入数字")
	fmt.Scanln(&a)

	for i := 1; i <= a ; i++ {//行
		for j := 1;j <= i ; j++ {//列
			fmt.Print("*")
		}
		fmt.Println()
	}
}
