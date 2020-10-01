package main

import "fmt"

func main() {
	fmt.Println(fun(6))
}
func fun(r float64)float64  {
	s:=(r*r)*3.14
	return s
}
