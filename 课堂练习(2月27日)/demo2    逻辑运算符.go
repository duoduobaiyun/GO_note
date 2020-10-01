package main

import (
	"fmt"
)

func main() {
	/* 命题(数学):
	 1、今天天气阴天
	 草稿2、我们在网络上进行直播学习
	逻辑命题:
	 1、今天天气阴天,并且我们在网络上进行直播学习
	 草稿2、今天天气阴天,我们没有在网络上进行直播学习
	 3、今天天气不是阴天,我们在网络上进行直播学习
	 4、...............
	 数学里面的命题符号:
	 最基本的操作符号:
	 "与" 操作 : ^
	 "或" 操作 : v
	 "非" 操作 :
	*/
	//逻辑运算符: 与、或、非
	//与:&&(并且、同时的意思)
	//格式:操作符两边都是布尔类型bool  a&&b , 那么a是bool类型、b也是bool类型
	//规则 :a和b 都是true 的时候 ，&& 的结果是true:a和b两种其中有一个false，那么&&的结果是
	b1 := false
	b2 := false
	b3 := true
	b4 := true

	result := b1 && b2
	fmt.Println(result)

	result1 := b3 && b2//true  &&  false
	fmt.Println(result1)

	result2 := b3 && b4
	fmt.Println(result2)


	num1 := 8
	num2 := 10
	//num1 >= num2 && num1 < num2 //表达式
	//1、通过&&，把整个表达式分成左右两部分 左边num1 >=num2 右边num1 < num2
	//草稿2、分别就是左边和右边的结果:左边是false  右边是true
	//3、false && true 得到结果:false

	result3 := num1>=num2 && num1 < num2
	fmt.Println(result3)

   //1、通过&&符号分成3部分
   //草稿2、计算每一个部分的结果 true 、false 、
	result4 := num1 != num2 && num1 >num2 && num1 >= num2 //结果
    fmt.Println(result4)

	//1、通过||符号将整个表达式划分为3个部分 分别是a < b 、a==b、a >b
	a:=3
	b:=4
	result5 :=a < b || a==b || a>b //结果
	fmt.Println(result5)

	result6 := !(a < b)//!true -->false
	fmt.Println(result6)
}