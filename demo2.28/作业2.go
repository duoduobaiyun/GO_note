package main

import "fmt"

func main() {
	b  := 2008
	res := b >> 0
	fmt.Println(res)

	res1 := b >> 1
	fmt.Println(res1)

	res2 := b >> 2
	fmt.Println(res2)

	res3 := b >> 3
	fmt.Println(res3)

	sun := res + res1 +res2 + res3
    fmt.Println(sun)
}
