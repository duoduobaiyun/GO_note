package main

import "fmt"

func main() {
n:=fun1(9,10,2,5)
fmt.Printf("%v",n)
}

func fun1(a,b,c,d int)*[4]int   {
	arr:=[4]int{a,b,c,d}
	for a:=1; a<len(arr); a++ {
		for b:=0;b<len(arr)-a ; b++ {
			if arr[b]>arr[b+1] {
				arr[b],arr[b+1]=arr[b+1],arr[b]
			}
		}
	}
	return &arr
}