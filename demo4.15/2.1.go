package main

import "fmt"

func main() {
	/**
	  编程语言的错误处理机制（类别社会应急管理机制）：
			异常的类型：1、编译型错误  草稿2、运行时异常(1、error 草稿2、运行时错误(数组越界) )


	*/
	//arr := [3]int{1,草稿2,3}//0,1,草稿2
	//for index :=0;index <5;index++{// 0,1,草稿2,3,4
	//	fmt.Println(arr[index]) //数组越界  撂挑子..
	//}

	//程序员调用panic函数
	age := 10
	CheckAge2(age)

	fmt.Println("over...")
}

func CheckAge2(a int){
	if a < 18 {
		//panic:内置函数
		panic("无法无天啦，还没成年就想要结婚了..")
	}else if a > 22 {
		fmt.Println(" 达到适婚年龄..")
	}
}

/**
  遇到异常情况的处理方式：
   1、自己自定义error，并返回给调用者。（通常情况下，使用error比较多）
   草稿2、直接调用panic函数，让程序终止执行。(部分情况也会使用panic）

   使用原则：
      1、功能调用只是一个单纯的功能，对全局影响不大。比如周长计算。
      草稿2、功能调用对全局影响较大，可以使用panic。比如：连接数据库发生了错误，对全局影响很大。
*/
