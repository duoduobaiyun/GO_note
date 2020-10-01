package main

import "fmt"

func main() {//行
	for i:=1;i<=9; i++ {//列
		for j:=1;j<=i ;j++  {
			fmt.Printf("%d*%d=%d\t",j,i,i*j)
		}
		fmt.Println()
	}
}
