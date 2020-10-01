package main

import (
	"fmt"
)

func main() {
	a:=make(map[string]string)
	a["姓名"]="刘星"
	a["性别"]="男"
	a["年龄"]="30"
	a["婚姻情况"]="已婚"
	a["资产"]="1亿"
	fmt.Println(a)

	b:=make(map[string]string)
	b["姓名"]="刘光"
	b["性别"]="男"
	b["年龄"]="18"
	b["婚姻情况"]="未婚"
	b["资产"]="8百万"
	fmt.Println(b)

	c:=make(map[string]string)
	c["姓名"]="刘雨欣"
	c["性别"]="女"
	c["年龄"]="22"
	c["婚姻情况"]="未婚"
	c["资产"]="2万"
	fmt.Println(c)

	s1:=make([]map[string]string,0,3)
	s1=append(s1,a)
	s1=append(s1,b)
	s1=append(s1,c)

	for key,val:=range s1{
		fmt.Printf("第%d个人的信息是\n",key+1)
		fmt.Printf("\t姓名:%s\n",val["姓名"])
		fmt.Printf("\t年龄:%s\n",val["年龄"])
		fmt.Printf("\t婚姻情况:%s\n",val["婚姻情况"])
		fmt.Printf("\t资产:%s\n",val["资产"])
	}
}
