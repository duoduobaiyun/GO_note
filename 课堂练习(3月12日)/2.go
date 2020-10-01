package main

import "fmt"

func main() {
		numArr := [...]int{20, 18, 15, 10, 5,2,-1}
		//目标：完成排序，从小到大的顺序进行排序
		//选择排序
		min := 0 // 设置一个变量, 作为最小值
		for index := 0;  index < len(numArr); index++  {
			//根据index，获取到numArr数组的元素，赋值给min变量，作为可能的最小值
			min = numArr[index]
			//index+1 是index的下一位
			for i := index +1; i < len(numArr); i++{// 内层for循环：
				//如果下一个位置上的元素 比 我们的min 要小，则要交换两者的位置
				if numArr[i] < min {
					numArr[index],numArr[i] = numArr[i],numArr[index]// 操作多元素的交换
					min = numArr[i]//把最小值设置成新找到的这个值
				}
			}
		}
		fmt.Println(numArr)
}