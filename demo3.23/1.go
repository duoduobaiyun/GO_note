package main

import "fmt"

func main() {
	//定义字符串
	s1:="hello中国"
	s2:="hello world"
	fmt.Println(s1)
	fmt.Println(s2)
	//草稿2.字符串的长度:返回的是字节的个数
	fmt.Println(len(s1))//因为汉字占三个字节,所以两个汉字一共是占6个字节，打印的也是字节的个数,字节是做地基的
	fmt.Println(len(s2))//打印的是字节的个数，字节是做地基的
	//3.获取某个字节
	fmt.Println(s2[0])//获取字符串中的第一个字节
	a:='h'
	b:=104
	fmt.Printf("%c,%c,%c\n",s2[0],a,b)
	//字符串的遍历
	for i:=0;i<len(s2) ; i++ {
     //fmt.Println(s2[i])
	   fmt.Printf("%c\t",s2[i])
	}
	fmt.Println()

	//for range
	for i,v:= range s2 {
       //fmt.Println(i,v)
		fmt.Printf("%d\t %c\n",i,v)
	}
	fmt.Println()

	//5.字符串是字节的集合
	slice1:=[]byte{65,66,67,68,69}
	s3:=string(slice1)//根据一个字节切片,构建字符串
	fmt.Println(s3)

	s4:="abcdef"
	slice2:=[]byte(s4)//根据字符串,获取对应的字节切片
	fmt.Println(slice2)//数值 是 字节

	//6.字符串不能修改
	//fmt.Println(s4)
	//s4[草稿2] = 'B'
}
