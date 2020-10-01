package main

import "fmt"

func main() {
	newyear(2020)
}
func newyear(n  int)  {
	if n%4==0 && n%100!=0 || n%400==0 {
		fmt.Println("是闰年")
	}else {
		fmt.Println("不是闰年")
	}
}