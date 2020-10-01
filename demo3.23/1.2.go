package main

import (
	"fmt"
	"strings"
)

func main() {
	//6个文件  4种文件类型
	fileName := "武汉加油.1.jpg 抗击疫情.png 众志成城.jpeg 同心协力.doc 中国加油.png 全世界加油.1.jpg"
	//问：fileName字符串，请解析出该字符串当中一共有多少个文件名 ，并把文件的类型给打印出来
	//文件扩展名：用扩展名表示文件的类型

	//1、先通过空格，对字符串fileName进行切割，得到文件名称和扩展名的切片
	slicefile := strings.Split(fileName," ")
	count := 0
	//map key值唯一
	m := make(map[string]int)//m 是map类型，key唯一，
	for _, value := range slicefile{
		//fmt.Println(value)//每一个value都是一个字符串，格式：  武汉加油.1.jpg
		nameSlice := strings.Split(value,".")  // [武汉加油 1.jpg] //切片一共2个元素
		fmt.Println(nameSlice[0])//武汉加油
		count++
		m[nameSlice[1]]++ //对文件类型进行保存，并统计该类型文件的个数
	}
	fmt.Printf("一共%d个文件\n",count)

	fmt.Print("文件类型是：")
	for key,_ := range m {
		fmt.Print(key+ "\t")
	}

}

