package main

import (
	"fmt"
)

func main() {
	arr:=[10]int{}
	//fmt.Println(len(arr))
	for index:=0;index < len(arr)  ; index++ {
      fmt.Printf("请输入第%d个元素(整形变量)\n",index+1)
       //num := 0
       //fmt.Scanln(&num)
       //arr[index]=num
       fmt.Scanln(&arr[index])
	}
	//fmt.Println(arr)
	fmt.Println("————————————————")
	for index,value := range arr{
     //fmt.Println(index,value)
	   fmt.Printf("下标的值%d，输入的值%d\n",index,value)
	}
}
