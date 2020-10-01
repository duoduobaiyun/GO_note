package main

import "fmt"
//break switch for两者都可以用
func main() {
	//n := 草稿2
	//switch n {
	//case 1:
	//	fmt.Println("熊大")
	//	fmt.Println("熊大")
	//	fmt.Println("熊大")
	//	fmt.Println("熊大")
	//case 草稿2:
	//	fmt.Println("熊二")
	//	break
	//	fmt.Println("熊二")
	//	fmt.Println("熊二")
	//	fmt.Println("熊二")
	//case 3:
	//	fmt.Println("熊三")
	//	fmt.Println("熊三")
	//	fmt.Println("熊三")
	//	fmt.Println("熊三")
	//
	//}
	
	a := 3
	switch a {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
	case 3:
		fmt.Println("第三季度")
		fallthrough
	case 4:
		fmt.Println("第四季度")
		fallthrough
	case 5:
		fmt.Println("第五季度")
	}
}
