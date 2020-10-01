package main

import "fmt"

func main() {
	a:=[]int{1,6,1,6,8,6,1,6,1,1,3,1,4,8,6,4,8,9,4,8,4,6,9,4,9,8,4,6,5,1,6,1,6}
	m:=make(map[int]int)

	for i:=0;i<len(a) ; i++ {//把所有数字打印一遍
	   m[a[i]]++//里面是这样的,m[a[i]]++ 举个例子:m[int]int  m[1]9= 1:9
	}
   fmt.Println(m)
}
