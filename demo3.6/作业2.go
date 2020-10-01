package main

import "fmt"

func main() {
	var b int
	fmt.Println("请输入数字")
	fmt.Scanln(&b)

	for i := 1; i <= b; i++ {
		for j := 1; j <= i; j++ {
			if j == 1 || i == b || i == j + 1{//j==1是指从左往右数第一列，，i==b是指第b行，，最上面的i和j已经执行过了，下面是指条件 i==j可能是 i==1 ==  j==1  两点连成一条直线
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
      //fmt.Println()
}