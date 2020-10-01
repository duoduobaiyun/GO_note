package main

import (
	"fmt"
)

func main() {
	arr1 := [4]int{1,2,3,4}
	arr2 := arr1
	arr2[1] = 5
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(&arr1[0]) //查看arr1的内存空间的地址
	fmt.Println(&arr2[0]) //查看arr2的内存空间的地址

	//切片
	s1 := make([]int,4)
	fmt.Println(s1,len(s1),cap(s1))//打印结果：[1 5 3 4]
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3
	s1[3] = 4

	// s2 := s1
	s2 := s1 //
	s2[1] = 5
	fmt.Println(s1,len(s1),cap(s1))//打印结果：[1 5 3 4]
	fmt.Println(s2,len(s1),cap(s1))//打印结果：[1 5 3 4]

	s3 := s2
	s3[3] = 9
	fmt.Println("s1 :",s1)
	fmt.Println("s2 : ",s2)
	fmt.Println("s3 :", s3)

	//打印地址
	fmt.Println(&s1[0])
	fmt.Println(&s2[0])
	fmt.Println(&s3[0])
}

