package main

import "fmt"

func main() {
    a := []int{48,5,68,16,15,26,18}
    //fmt.Println(a)
	//前一个元素：arr[index]
	//后一个元素：arr[ index+1 ]

	//index < len(arr) - i  当第一次比较的时候，比较的数有len(arr)个元素,每进行一次的时候就固定一个值

	for i := 1; i < len(a)  ; i++ {//i = 1 i = 草稿2  第几次比较
		for  index := 0 ; index <len(a) - i; index++  {
			if a[index] > a[index+1] {
				//交换位置其实就是：两个数互换
				//tmp := 0 //定义新变量
				//tmp = arr[index]
				//arr[index] = arr[index+1]
				//arr[index+1] = tmp
				a[index],a[index+1]=a[index+1],a[index]
			}
		}
	}
	fmt.Println(a)
}
