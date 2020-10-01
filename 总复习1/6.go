package main

import "fmt"

func main() {
	/**
	  定义一个长度为6的字符串数组，并进行元素的初始化赋值；
	  要求：将字符串数组中各元素的地址保存到一个数组中，并打印该指针数组的元素值。
	   最后通过操作指针数组元素修改 原字符串数组中最后一个元素的元素值
	*/

	arry:=[6]string{"zhong","guo","jia","you","wang","sui"}
	arr1:=[6]*string{}//指针数组
	for i:=0;i<len(arr1) ; i++ {
		arr1[i]=&arry[i]
	}
	fmt.Println(arr1)

	*arr1[5]="le"
	fmt.Println(&arr1)
}
