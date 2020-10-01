package main

import "fmt"

func main() {
	var a [20]int = [20]int{3,9,10,11,8,4,9,1,20,10,11,21,19,3,8,4,1,10,20,21}
	arr1 :=make(map[int]int)
	b:=0
	for i:=0;i<len(a);i++ {
		b = a[i]
		if arr1[b] != 0 {
			continue
		}
		for j := 0; j < len(a); j++ {
			if b == a[j] {
				arr1[b]++
			}
		}
	}
	fmt.Println(arr1)
}
