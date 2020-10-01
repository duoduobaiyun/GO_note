package main

import "fmt"

func main() {
	/**
	  有一个数组，如下，要求统计出数组当中各个元素出现的个数。
	*/
	numArr := [20]int{3,9,10,11,8,4,9,1,20,10,11,21,19,3,8,4,1,10,20,21}

	map1 := make(map[int]int)
	for index := 0; index < len(numArr); index++{//获取数组的每一个元素
		map1[numArr[index]]++
	}
	fmt.Println(map1)

	/**
	  1、分析思路（分几步走）
	  草稿2、程序是一点一点写出来。先写大的步骤，再完善细节代码逻辑
	*/
}
