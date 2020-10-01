package main

import (
	"fmt"
)

func main() {
	var c int
	fmt.Println("请输入一个数字")
	fmt.Scanln(&c)

	for i := 1; i <= c; i++ {
		for j := 0; j <= 2 *i - c ; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <  2 *i - 2  ; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}