package main

import (
	"fmt"

)

func main() {
	//switch 和 if 都只有一个初始化语句
	//var  language string
	//fmt.Printf("请输入编程语言:%s\n",language)
	////fmt.Scanln(&language)
	//
	//switch   fmt.Scanln(&language); language {
	//case "python":
	//	fmt.Println("个人感觉挺好的")
	//case "ruby":
	//	fmt.Println("只是听过")
	//case "go":
	//	fmt.Println("目前正在学，感觉容易上手")
	//default:
	//	fmt.Println("您也真是的，不要乱猜，编程的艺术你不懂")
	//}
     fmt.Println("请输入一个数字")
	var a int
	if   fmt.Scanln(&a) ; a > 10 {
          fmt.Println("正确")
	}
}
