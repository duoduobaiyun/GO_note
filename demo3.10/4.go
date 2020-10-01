package main

import "fmt"

func main() {
	a :=[10]int{}
	//for _,value :=range a {
		for i := 0; i<10 ; i++ {
			fmt.Scanln(&a[i])
			fmt.Println(a)
		//}
	}
	//fmt.Println(len(a))
}

