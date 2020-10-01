
package main

import (
	"fmt"
)

func main() {
	//作业一：定义圆形结构体，包含有半径、周长、面积等属性；
	//定义圆形结构体所拥有的两个方法：计算面积、计算周长。在main函数中实现调用。

	a:=Circle{
		radius:    5,
		perimeter: 12,
		area:      100,
	}
	a.fun1()

}

type Circle struct {
	radius int
	perimeter  int
	area  int
}

func (c Circle) fun1()  {
	fmt.Printf("面积:%d,周长:%d",c.perimeter,c.area)
}