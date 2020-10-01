package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "==yuhongwe=="
	fmt.Println(str)
	//去除字符串 开头和末尾 的指定的内容（掐头去尾留中间）
	s := strings.Trim(str,"==")
	fmt.Println(s)
	//对于去除字符串当中开头和末尾的空格，有一个专门的方法
	str1 := "  yuhongwe  "
	s2 := strings.TrimSpace(str1)
	fmt.Println(s2)
}