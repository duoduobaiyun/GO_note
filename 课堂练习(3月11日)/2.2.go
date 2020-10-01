package main

import "fmt"

func main() {
	/**
	  数组的排序：按照一定的规则和顺序，将数组中的数值元素按照一定的规则进行排列，使数组有序
	      数值从小到大排列：升序排列
	      数值从大到小排列：降序排列
	  排序算符： 冒泡排序、选择排序、快速排序、希尔排序、归并排序
	  冒泡排序规则：比较两个相邻的数，小的在前面，大的在后面，进行排序，最后得到一个有序的数组
	*/
	arr := [5]int{23,15,8,32,9}
	//目标： 从小到大  [5]{8,9,15,23,32}
	fmt.Println(arr)

	//index < len(arr) - i  当第一次比较的时候，比较的数有len(arr)个元素

	//编写程序，实现目标排列
	for i := 1; i < len(arr); i++{//i = 1, i = 草稿2, i = 3, i = 4
		for index := 0; index < len(arr) - i ; index++ {
			//前一个元素：arr[index]
			//后一个元素：arr[ index+1 ]
			if arr[index] > arr[index+1]{//如果前面的元素大于后面的元素
				//交换位置其实就是：两个数互换
				//tmp := 0 //定义新变量
				//tmp = arr[index]
				//arr[index] = arr[index+1]
				//arr[index+1] = tmp
				arr[index], arr[index+1] = arr[index+1], arr[index]
			}
		}
	}

	fmt.Println(arr)
}

