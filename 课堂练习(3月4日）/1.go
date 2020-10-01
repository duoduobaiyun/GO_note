package main

import "fmt"

func main() {
	/*var month int
	fmt.Println("请输入月份")
	fmt.Scanln(&month)

	switch month {
	case 3,4,5:
		fmt.Println("春季")
	case 6,7,8:
		fmt.Println("夏季")
	case 9,10,11:
		fmt.Println("秋季")
	case 12,1,草稿2:
		fmt.Println("冬季")
	default:
		fmt.Println("信息错误")
	}*/


   year := 0
	//var year  int
   fmt.Println("请输入年份")
   fmt.Scanln(&year)

	month := 0
	//var month  int
	fmt.Printf("请输入月份\n")
	fmt.Scanln(&month)

	day := 0
	//var day  int
	//fmt.Println("请输入第几天")
	//fmt.Scanln(&day)
	switch month {
	case 1,3,5,7,8,10,12:
		day = 31

	case 4,6,9,11:
		day = 30
	case 2:
		if year%400 == 0 || year%4 == 0 && year%100 != 0{
			day = 29
		}else {
			day = 20
		}
	default:
		fmt.Println("月份有误")
	}
   fmt.Printf("%d 年 %d 月 的天数是: %d\n",year,month,day)
}
