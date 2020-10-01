package main

import "fmt"

func main() {

	a := []int{5, 9, 20, 120, 450, 340}

	for i := 1; i < len(a); i++ {
		for index := 0; index < len(a)-i; index++ {
			if a[index+1] > a[index] {
				a[index+1], a[index] = a[index],a[index+1]
			}
		}
		fmt.Println(a)
	}
}