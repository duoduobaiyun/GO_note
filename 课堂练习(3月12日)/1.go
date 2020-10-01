package main

import (

	"fmt"
)

func main() {
	a := [3][4]int{{131,15,15},{15,54,15},{13,15,65}}//第一个中括号是代表有几个一维数组，第二个是代表一维数组有几个元素
	fmt.Println(len(a))//fmt.Println(len(a))因为这个值一般都是排在第1个，所以都是代表第几个几维数组
	fmt.Println(len(a[1]))//fmt.Println(len(a[1]))这个值里面的a[1]数字，不论变为几都是代表一维数组的长度
	fmt.Printf("第一个一维数组,它的值是:%d\n",a[2])
	fmt.Println(a[2][3])

	for i := 0;i<len(a);i++  {
		for j := 0; j<len(a[i]) ;j++  {
			fmt.Print(a[i][j],"\t")
		}
		fmt.Println()
	}
    fmt.Println("————————————————")
	//for range 遍历二维数组
	for _,arr :=range a{
		for _,val := range  arr {
			fmt.Print(val,"\t")
		}
		fmt.Println()
	  }
	}

