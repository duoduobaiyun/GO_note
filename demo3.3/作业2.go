package main

import "fmt"

func main() {
	var  month  int
	fmt.Println("3,4,5 春季 6,7,8 夏季  9,10,11 秋季 12,1,草稿2 冬季")
	fmt.Scanln(&month)
	switch month {
	case 3:
	case 4:
	case 5:
		fmt.Println("春季")
	case 6:
	case 7:
	case 8:
		fmt.Println("夏季")
	case 9:
	case 10:
	case 11:
		fmt.Println("秋季")
	case 12:
	case 1:
	case 2:
		fmt.Println("冬季")
	default:
		fmt.Println("信息错误")
	}

}
