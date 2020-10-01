package main

import "fmt"

func main() {
	a := 30
	b := 60
	c := 20
	fmt.Printf("a的十进制数:%d\n",a)
	fmt.Printf("b的十六进制数:%X\n",b)
	fmt.Printf("c的八进制数:%o\n",c)

    res1 :=  a & b
    fmt.Print(res1)
}
