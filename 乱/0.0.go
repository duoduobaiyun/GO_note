package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("请输入数字")
	fmt.Scanln(&a)

	for i := 0; i <= a ; i++ {
		for j := 0; j <  i ; j++ {
			fmt.Print("*")
		}
		for k := 1; k < i ; k++ {
			if k==1 || k==i  || k==a {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}
}