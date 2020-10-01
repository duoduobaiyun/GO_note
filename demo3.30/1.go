package main

import "fmt"

func main() {

	fmt.Println(fun1(3))
}

func fun1(n  int)int  {
	if n ==1 {
		return 1
	}
	return fun1(n-1)*10+n
}