package main

import "fmt"

func main() {
   var  week int
   fmt.Println("如果为1，就输出星期一，如果为2，就输出星期二，如果为3，就输出星期三，如果为4，就输出星期四，如果为5，就输出星期五，如果为6，就输出星期六，如果为7，就输出星期天")
   fmt.Scanln(&week)
	if week <= 7  {
		if week == 1{
		    fmt.Println("星期一")
		}else if week == 2 {
			fmt.Println("星期二")
		}else if week==3 {
			fmt.Println("星期三")
		}else if week==4 {
			fmt.Println("星期四")
		}else if week==5 {
			fmt.Println("星期五")
		}else if week==6 {
			fmt.Println("星期六")
		}else if week==7 {
			fmt.Println("星期天")
		}

	} else {
		fmt.Println("错误信息")
	}

}
