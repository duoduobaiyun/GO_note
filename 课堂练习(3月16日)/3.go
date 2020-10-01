package main

import "fmt"

func main() {
	/**
	  删除切片slice的第二个元素值
	*/
	slice := []int{2,5,7,9,11,21,45}
	fmt.Println(slice)

	//第一部分
	subSlice1 := slice[:1] //从头开始切，开头不写
	fmt.Println(subSlice1)

	//第二部分
	subSlice2 := slice[2:]//切刀最后，结束值可以不写
	fmt.Println(subSlice2)

	//将第一部分和 第二部分进行拼接
	slice3 := append(subSlice1,subSlice2...)
	fmt.Println(slice3)

	//两个切片直接拼接
	//1、直接拼接两个切片
	s1 := []int{1,3,5}
	s2 := []int{2,4,6}
	fmt.Println(s1,s2)
	//s3 := append(s1,s2...)
	//fmt.Println(s3)


	s3 := make([]int,3)
	//草稿2、挨个添加数据
	for i := 0; i< len(s2); i++{
		s3 = append(s1,s2[i])
	}
	fmt.Println(s3)

}