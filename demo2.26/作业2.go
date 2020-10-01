package main

import "fmt"

func main() {
	name := "wuhanjiayou,zhongguojiayou"
    fmt.Println(len(name))

	wuhan := name[0:5]
	fmt.Println(wuhan)

	zhonguo := name[12:20]
	fmt.Println(zhonguo)

	wuhanjiayou := name[0:11]
	fmt.Println(wuhanjiayou)

	zhongguojiayou := name[12:26]
	fmt.Println(zhongguojiayou)
}
