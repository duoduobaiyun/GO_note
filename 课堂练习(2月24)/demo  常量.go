package main

import "fmt"

func main(){
	//数学知识:PAI 是一个固定数值:3.1414926 -3.1415927,往往设置
	//格式一:const 常量名字 数据类型 = 内容值
	const  PAI float32  = 3.14//定义了一个常量,名字为PAI、
	fmt.Println(PAI)
	//PAI = 3.1415926

	const BAIDU = "http://www.baidu.com"//定义个常量，名字为BAIDU,具体的值为。。。。
	fmt.Println(BAIDU)

	//规范:常量值在定义的时候，其名字往往会使用全部大写来进行和变量区分

	const(
		num = 100
		name = "LIUKUN"
		sex = "男"
	)
	//常量的举例
	const (
		MONDAY = "星期一"
		TUESDAY = "星期二"
		WEDNESDAY = "星期三"
		THURSDAY = "星期四"
	)
	//四季:春夏秋冬  、 一年十二月
	//常量组
	const (
		x int = 10 //x常量， int类型，数值10
		y int = 10 //y常量， int类型，数值10
		z int = 10 //z常量， int类型，数值10
		name1 string = "liukun"
		sex1 string = "男"
	)
 fmt.Println("x,y,z",x,y,z)
}
