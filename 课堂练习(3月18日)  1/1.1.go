package main

import "fmt"

func main() {
	a:=[20]int{3,9,10,11,8,4,9,1,20,10,11,21,19,3,8,4,1,10,20,21}
	b:=make(map[int]int)

	for i:=0;i<len(a) ; i++ {
		b[a[i]]++//key值固定不变,a是key值,i是每个元素的值
	}
	fmt.Println(b)
}