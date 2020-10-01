package main

import "fmt"

func main() {
	var a  int
	fmt.Println("请输入数字")
	fmt.Scanln(&a)

	for i := 0; i <= a ; i++ {
		for j := 0; j < a - i ; j++ {
           fmt.Print(" ")
		}
		for k := 0; k < 2 * a - 1 ; k++ {
          fmt.Print("*")
		}
		fmt.Println()
	}
}
