package main

import "fmt"

func main() {
	//第一种方法,把字符串转化为数组来算,利用数组的特性，和map的特性
	/*a:="zhongguojiayouwuhangjiayou"
	b:=[35]string{}
	for i:=0;i<len(a) ; i++ {//获取到每一个字母，放到数组当中
		b[i]=a[i:i+1]
	}

	m:=make(map[string]int)
	count:=0count变量用于获取该字母的次数
	for j:=0;j<len(a) ; j++ {//拿到数组中的每个元素
		key:=b[j]//因为k是唯一的，所以不能重复,右边的值赋值给左边的
		count=m[key] //因为count是次数,又因为右边的值是是map,特性:key唯一不变,所以它的意思是,遇到相同的数,计数加1
		count++//对count值进行++，增加1
		m[key]=count//左边key固定不变,右边value值
	}

	for key,val :=range m {
		fmt.Println(key,val)
	}
	 */

	a:="zhongguojiayouwuhangjiayou"
	m:=make(map[string]int)
	for i:=0 ;i<len(a) ; i++ {
		m[a[i:i+1]]++
	}
	fmt.Println(m)


}