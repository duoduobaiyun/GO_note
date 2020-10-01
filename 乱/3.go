package main

import "fmt"

func main() {
	var n int
	fmt.Println("请输入一个数")
	fmt.Scanln(&n)

	for  i := 1; i <= n  ; i++ {//行   因为i 执行结果是 ( 1，草稿2 ，3，4，5  )  所以就有了k这个变量
		for j := 0; j < i -1; j++ {//列
			fmt.Print("  ")
		}
		for k := 0; k < 2 * i -1 ; k++ {// 草稿2  *  (1,草稿2,3,4,5) - 1   因为
			fmt.Print("*")
		}
		fmt.Println()
	}

}
