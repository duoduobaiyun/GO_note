package main

import (
	"fmt"
)

func main() {
	//3位数的时候
	var num int
	for num = 0; num < 1000; num++ {
		if num >= 100 && num < 1000 {
			//个位数、十位数、百位数
			c := num % 10//个位
			d := num / 10 % 10//十位
			e := num / 100//百位
			if c*c*c+d*d*d+e*e*e == num {
				fmt.Println(num)
			}
		}
	}
}