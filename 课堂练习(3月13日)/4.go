package main

import "fmt"

func main() {
	s1 := make([]int,5,7)
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3
	s1[3] = 4
	s1[4] = 5
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}
