package main

import "fmt"

func main() {
	a :=[...]int{28,15,63,89,5,8}
	for i :=1;i < len(a);i++  {
		for index := 1 ; index <= len(a) - i; index++ {
			if a[index] > a[index+1] {
				a[index],a[index+1]=a[index+1],a[index]
			}
		}
		fmt.Println(a)
	}
	//fmt.Println(a)
}