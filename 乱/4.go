package main

import "fmt"

func main() {
	var   a  int
	fmt.Println("请输入数字")
	fmt.Scanln(&a)

	 out:for i := 1; i <=a  ; i++ {
		for j :=1; j<=a  ; j++ {
			if j == 3 {
				break out
			}
			fmt.Printf("i的值:%d,j的值:%d\n",i,j)
		}
	}
	fmt.Println("哈哈")
}
