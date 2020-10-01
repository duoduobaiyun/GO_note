package main

import "fmt"

func main() {
	var b int
	fmt.Println("请输入三角形高度")
	fmt.Scanln(&b)
	for i := 1; i < b ; i++ {
		for j := 1; j < i - 1 +b ; j++  {
			if j == b - i + 1 || j == b + i -1 {
				fmt.Print("*")
			}else {
				fmt.Print("")
			}
			fmt.Println()
		}
		for i = 1; i < 2 * b - 1 ; i++  {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
