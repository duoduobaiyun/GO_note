package main

import "fmt"

func main() {
	a:=[]int{1,2,3,4}
	b:=a
	fmt.Println(a)
	fmt.Printf("%p\n",a)//切片的地址
	fmt.Printf("%p\n",&a)//a的地址
	fmt.Printf("%p\n",b)//切片的地址
    fmt.Printf("%p\n",&b)//b本身的地址


}
