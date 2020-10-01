package main

import "fmt"

//打印58到23的数
//求1到100的值
//打印1-100内，能够被3整除，但是不能被5整除的数字，统计被打印的数字的个数，每行打印5个
func main() {
	//for i:= 58 ; i >= 23; i-- {
	//	fmt.Println(i)
	//}

	//for sun:=1;  sun <= 100; sun++ {
	//	fmt.Println(sun)
	//}

	for a:= 1; a <= 100 ; a++ {
		//fmt.Println(a)
		if a % 3 == 0 &&  a % 5 != 0 {
			fmt.Println(a)
		}
	}
}
