package main

import "fmt"

func main() {
	house:="Malibu  Point  10880, 90265"
	ptr:=&house
	fmt.Printf("ptr type := %T\n",ptr)
	fmt.Printf("adress:=%p\n",ptr)
	val:=*ptr
	fmt.Printf("val type:=%T\n",val)
	fmt.Printf("val:%s\n",val)
}
func swap(a,b *int)  {
	//取a指针的值,赋给临时变量t
	t:=*a
	//取b指针的值,赋给a指针指向的变量
	*a=*b
	//将a指针的值赋给b指针指向的变量
	*b=t
}

func main1()  {
	//准备两个变量,赋值1和2
	x,y:=1,2
	//交换变量值
	swap(&x,&y)
	//输出变量值
	fmt.Println(x,y)
}
