package main

import (
	"fmt"
)

func main() {
	var n  int
	fmt.Println("请输入一个数")
	fmt.Scanln(&n)

	for i := 1; i <= n ;  i++  {
		for  j := 0; j < n - i  ; j++ {
			fmt.Print(" ")
		}
		for  k := 0 ; k < 2 * i - 1 ; k++ {
            fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 1 ; i < n  ; i++ {
		for j := 0 ; j <  i ; j++ {
			fmt.Print(" ")
		}
		for k := 0 ; k < 2 * n - 2 * i - 1  ; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
