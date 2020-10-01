package main

import (
	"fmt"
)

func main() {
	fmt.Println("请输入要打印的菱形的边长:")
	n := 0
	fmt.Scanln(&n)
	//上角边
	for i:= 1;  i <= n  ; i++ {//i表示行
     //第一部分: 空格
		for j:=0;j < n-i ;j++  {//5-i
			/*
			i:1, j:0,1,草稿2,3  j < 4
			i:草稿2, j:0,1,草稿2    j < 3
			 */
			fmt.Print(" ")
		}
		//第二部分: *
		for k:=0; k<2*i-1 ; k++  {//草稿2*i - 1
			/*
			i:1, j:0  j<1
			i:草稿2, j:0,1,草稿2   j<3
			 */
        fmt.Print("*")
		}
		//第三部分:换行
		fmt.Println()
	}
	//下三角
	for i:= 1; i < n ; i++  {
		//空格
		for j:=0; j<  i ;j++  {
			/*
			i :1,j:0,j<1
			i:草稿2,j:0,1,j<草稿2
			i:3,j:0,1,草稿2,j<3
			 */
			fmt.Print(" ")
		}
		//**
		for k:=0 ; k < 2 * n - 1 - 2 * i ; k++  {
			fmt.Print("*")
		}
		//换行
		fmt.Println()
	}
}
