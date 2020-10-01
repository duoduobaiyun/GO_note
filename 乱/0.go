package main

import "fmt"

func main() {
	// for: 循环的去执行一段程序
	/**
	  Hello	Hello	Hello	Hello	Hello
	  Hello	Hello	Hello	Hello	Hello
	  Hello	Hello	Hello	Hello	Hello
	  Hello	Hello	Hello	Hello	Hello
	  Hello	Hello	Hello	Hello	Hello
	*/
	for num := 0; num < 3; num ++ {//行
		for i := 0; i < 3; i++{// 列
			fmt.Print("Hello\t")// 5个hello world
		}
		fmt.Println()
	}
	//多层for循环实现功能

}