package main

import "fmt"

func main() {
	var num  int
	fmt.Println("请输入数字")
	fmt.Scanln(&num)

	for i := 1; i <= num ; i++  {
		//fmt.Print(i)
		for j := -2; j < 2*num-2*i-1 ; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2 * i -1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

