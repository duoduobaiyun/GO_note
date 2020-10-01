package main

import "fmt"

func main() {
	a := []int{12,23,56,5,1,0,15,12,5}

  min := 0
	for index :=0 ;index < len(a) ;index++ {
		min=a[index]
		for i := index+1;i < len(a)  ;i++ {
			if a[i] < min{
				a[index],a[i]=a[i],a[index]
                min = a[i]
			}
		}
	}
	fmt.Println(a)

}