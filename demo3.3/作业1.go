package main

import "fmt"

func main() {
	var week  int  
	fmt.Println("如果为1，就输出星期一，如果为2，就输出星期二，如果为3，就输出星期三，如果为4，就输出星期四，如果为5，就输出星期五，如果为6，就输出星期六，如果为7，就输出星期天")
	fmt.Scanln(&week)
	switch week {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期天")
	default:
		fmt.Println("错误信息")
	}
}
