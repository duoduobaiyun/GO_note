package main

import "fmt"

func main() {

	/**
	  回文数、回文字符串： 诸如 aba, 1221,abcddcba等类似形式的，左右对称的数据，称为回文。
	  回文规则：  以字符串或者数字等给定的数据，从中间进行划分，左右对称，则称为回文
	*/
	str := "ABA"// 判断该字符串是否是回文数

	//1、先定位中间位置
	center := len(str) /2

	//草稿2、判断从中间位置开始，相等距离的左右两个位置上的元素是否相等
	preStr := ""
	lasStr := ""

	for index := 0; index <= center; index++{
		//拿到前部分的元素
		preStr = string(str[index])
		//拿到后半部分的元素
		lasStr = string(str[len(str) - 1 - index])

		//3、根据判断结果，进行下一步执行
		if preStr == lasStr {//如果两个元素值一致
			// 3.1 如果距离中间位置相同距离的元素值一样，判断下一位，直到全部判断完成
			//进行下一次比较
			if index  == center {
				fmt.Println("该字符串是回文数")
			}
			continue
		}

		// 3.草稿2 如果距离中间位置相同距离的元素值不一样，直接返回结论：不是回文。
		//如果两者不相等,直接跳出循环
		fmt.Println("该字符串不是回文字符串")
		break
	}

}

