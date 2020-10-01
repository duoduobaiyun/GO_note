package main

import "fmt"

func main() {
	a:=make([]int,4,4)
	a[0]=1
	a[1]=2
	a[2]=3
	a[3]=4

	fmt.Printf("%p\n",a)//地址:0xc000010340
	fmt.Printf("%p\n",&a[0])//地址:0xc000010340
	//总结:切片的地址跟第一个下标值地址一样
	fmt.Printf("%p\n",&a[1])//地址:0xc000010348
	fmt.Printf("%p\n",&a[2])//地址:0xc000010350
	fmt.Printf("%p\n",&a[3])//地址:0xc000010358
	//每一个元素的地址都是不同的
}
