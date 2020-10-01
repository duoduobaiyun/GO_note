package main

import "fmt"

func main() {
	fmt.Printf("作业二\n")
	var in [20]int = [20]int{3,9,10,11,8,4,9,1,20,10,11,21,19,3,8,4,1,10,20,21}
	fmt.Printf("%d\n",in)
	arr2 :=make(map[int]int)
	var b int=0
	for i:=0;i<len(in);i++ {
		b = in[i]
		if arr2[b] != 0 {
			continue
		}
		for j := 0; j < len(in); j++ {
			if b == in[j] {
				arr2[b]++
			}
		}
	}
}
