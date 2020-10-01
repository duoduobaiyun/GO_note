package main

import "fmt"

func main() {
   a:=[6]int{5,9,20,120,450,340}

	for index := len(a)-1; index>=0; index-- {
		for j :=index-1; j>=0; j-- {
			max:=a[index]
			if max >a[j]  {
				a[index],a[j]=a[j],a[index]
				}
		    }
		}
		fmt.Println(a)
	}



