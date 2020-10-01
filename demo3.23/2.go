package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	strings包下的关于字符串的函数
	 */
	s1:="helloworld"
	//1.是否包含指定的内容————>bool
	fmt.Println(strings.Contains(s1,"h"))
	//草稿2.是否包含chars中的任意的一个字符即可,只要有一个字节相同即可
	fmt.Println(strings.ContainsAny(s1,"abcd"))
	//3.统计substr在s中出现的次数!!! 次数!!!
	fmt.Println(strings.Count(s1,"lloo"))
	//4.以xxx的缀开头,xxx后缀结尾
	s2:="20190525课堂笔记.txt"
	if strings.HasPrefix(s2,"201905") {
		fmt.Println("19年5月的文件。。")
	}
	if strings.HasSuffix(s2,".txt") {
		fmt.Println("文本文档。。")
	}
	//helloworld
	fmt.Println(strings.Index(s1,"lloo"))//查找substr在s中的位置,如果不存在就返回-1
	fmt.Println(strings.IndexAny(s1,"abcdh"))//查找chars中任意的一个字符,出现在s中的位置
	fmt.Println(strings.LastIndex(s1,"l"))//查找substr在s中最后一次出现的位置，从后往前找
	fmt.Println(strings.IndexByte(s1,104))//查找指定的字节,以下标的方式出现，输入ascall码
	fmt.Println(strings.IndexRune(s1,'o'))//查找指定的类型，以下标的方式出现，在单引号输入s1中的字符

	//字符串的拼接
	ss1:=[]string{"abc","world","hello","ruby"}
	s3:=strings.Join(ss1,"-")
	fmt.Println(s3)

	//切割
	s4:="135 1616 61 8616"
	//ss2:=strings.Split(s4,"1")
	//fmt.Println(ss2)
	ss3:=strings.Split(s4," ")
	//fmt.Println(ss3)
	for i:=0;i<len(ss3) ; i++ {
		fmt.Println(ss3[i])
	}

	str := "WU HAN JIA YOU ZHONG GUO JIA YOU I LOVE CHINA"
	rs := strings.Split(str," ")
	fmt.Println(rs)

	//指定切割多少次
	rs1 :=strings.SplitN(str," ", 3)//n指的切割后的个数。全切：n = -1
	for _, value := range rs1 {
		fmt.Println(value)
	}


    //自己拼接自己count次
    s5:=strings.Repeat("hello ",5)
    fmt.Println(s5)

    //替换
    //helloworld
    s6:=strings.Replace(s1,"l","*",2)//如果n = -1,又因为-1是无效数字,所以系统默认全部选中
    fmt.Println(s6)

	str1 := "WU HAN JIA YOU ZHONG GUO JIA YOU I LOVE CHINA"
	//把字符串当中JIAYOU，替换成WEIWU
	//all英文单词：所有的，全部  replace
	st := strings.ReplaceAll(str1,"JIA YOU","WEI WU")
	fmt.Println(st)

    s7:="hello world**123.."
    fmt.Println(strings.ToUpper(s7))//小写字母转换为大写字符
	fmt.Println(strings.ToLower(s7))//大写字母转换为小写字符

	/*
	substring(start,end)————>substr
	str[start:end]——————>substr
	    包含start,不包含end下标
	 */
	fmt.Println(s1)
	s8:=s1[:5]
	fmt.Println(s8)
	fmt.Println(s1[5:])
}
