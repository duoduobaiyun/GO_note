package main

import "fmt"

func main() {
	var a int
	fmt.Println("请输入数字")
	fmt.Scanln(&a)

	var count  = 0
	for i := 1 ; i < a ; i++ {
	for j := 1; j < i ; j++ {
		if i%j==0 {
			count++
			fmt.Print(j)
		    }
		}
	}
}
