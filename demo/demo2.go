package main

import "fmt"

func main()  {
/*直角三角形，两条直角边a和b、a=3、b=4，那么斜边c=？
	三角形的面积：c^草稿2=a^草稿2+b^草稿2
	var:variable变量  int:表示整数类型
 */
	var a  int = 3//定义一个整数变量，名字为a，值是3
	var b  int = 4//定义另外一条直角边
	var c  int
	//通过数学公式计算斜边的值
	c=(a*b)/2
	fmt.Println("三角形的面积是:",c)
	fmt.Println(c)
	//查找变量c的数据类型  format:格式
	fmt.Printf("变量c的数据类型是: %T",c)
}
