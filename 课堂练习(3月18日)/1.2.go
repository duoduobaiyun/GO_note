package main

import (
	"fmt"
	"strings"
)

func main() {
	//给定字符串，打印出字符串当中每个字母出现的次数
	str := "WUHANJIAYOUZHONGGUOJIAYOUILOVECHINA"
	fmt.Println(str)

	//方法一：
	//str[x:y] 通过下标获取某一段字符串
	s := ""
	m := make(map[string]int) // string表示每个字母， int表示字母出现的个数
	//map的特点：key唯一， value值可修改
	for index := 0;  index < len(str); index++{//遍历每一个字母  n
		s = str[index:index+1]//得到每一个字符串中的字母
		if m[s] > 0{// 如果m[s] >0 ，则表示map当中已经有该字母的次数统计了，也就是说已经统计过该字母了，跳过该字母
			continue //跳过本次循环，继续执行下一次循环
		}//总的来说就是把每种不同元素打印出来
		m[s] = 1// 先赋值为1 m["w"] = 1  m["U"] = 1
		for i := index+1; i < len(str); i++{//从遇到的字母的下一位开始，遍历到字符串的末尾，统计字母出现的个数 n-1
			if str[i:i+1] == s{//出现了一个相同的字母，计数要+1
				m[s]++
			}
		}
		fmt.Println(m)
	}

	//总的循环次数： n * (n -1) ~~ n^草稿2 如果n这个数很大 。
	//在执行程序过程当中，程序的性能：花的时间少，干更多的活。 效率高

	//第二种解决方案
	str1 := "WUHANJIAYOUZHONGGUOJIAYOUILOVECHINA"
	// 数组
	//1、将字符串变成一个数组 [35]string{"W","U","H","A","N","J",.....}
	//草稿2、利用map数据结构本身的特点进行统计
	//   map数据结构：1、key唯一、与value对应
	//               草稿2、map[key] = value1、value2 如果key已经存在，value会修改成新的值
	//               3、如果不存在key值，则map[key]返回一个默认值 int类型的默认值是0
	strArr := [35]string{}
	for index := 0; index < 35; index++{//获取到每一个字母，放到数组当中
		strArr[index] = str1[index:index+1]
	}

	map1 := make(map[string]int)// value = map[key]
	count :=0 //count变量用于获取该字母的次数
	for index := 0; index < len(strArr); index++{//拿到数组中的每个元素
		key := strArr[index]//因为k是唯一的，所以不能重复,右边的值赋值给左边的
		count = map1[key] //因为count是次数,又因为右边的值是是map,特性:key唯一不变,所以它的意思是,遇到相同的数,计数加1
		count++ //对count值进行++，增加1,由于++后,值发生了改变
		map1[key] = count//count++后,需要重新赋值给map1[key]
	}

	//遍历输出
	for key,value := range map1{
		fmt.Println(key,value)
	}

	// 执行时间： 转数组 n  +  遍历数组 n  +   遍历输出 n


	//第三种解决方案
	//对字符串转数组，进行优化
	// "YU HONG WEI" 按照空格进行切割
	//字符串转切片的方法: 字符串切割
	/**
	  字符串的切割：按照一定的规则，将一个字符串划分成几个部分
	  例如：str := "WU HAN JIA YOU, ZHONG GUO JIA YOU, I LOVE CHINA"
	    要求：使用“,"对字符串进行切割
	      1、第一个逗号：① WU HAN JIA YOU   ② ZHONG GUO JIA YOU, I LOVE CHINA
	      草稿2、第二个逗号：① WU HAN JIA YOU   ② ZHONG GUO JIA YOU  ③ I LOVE CHINA

	    str2 := "WUHANJIAYOUZHONGGUOJIAYOUILOVECHINA" 切割字符串,切割成单个字母

	  go语言当中的字符串切割： strings.Split
	    split: 切割,把....分开，切开
	*/
	//包:fmt rand  time strings字符串
	str2 := "WUHANJIAYOUZHONGGUOJIAYOUILOVECHINA"
	strs := strings.Split(str2,"")//使用空字符串将字符串str2进行切割，切割成单个的字母, 返回的是一个切片
  fmt.Println(strs)
	map2 := make(map[string]int)// value = map[key]
	for index := 0; index < len(strs); index++{//拿到数组中的每个元素
		map2[strs[index]]++//就是下面的简写
		//count :=0 //count变量用于获取该字母的次数
		//key := strArr[index]//因为k是唯一的，所以不能重复,右边的值赋值给左边的
		//count = map1[key] //因为count是次数,又因为右边的值是是map,特性:key唯一不变,所以它的意思是,遇到相同的数,计数加1
		//count++ //对count值进行++，增加1,由于++后,值发生了改变
		//map1[key] = count//count++后,需要重新赋值给map1[key]
	}
	fmt.Println(map2)


}

