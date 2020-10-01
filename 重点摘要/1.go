package main

import "fmt"

func main() {
	/**
	  总结：结构体
	    1、描述性信息：往往是一些名词,比如说：name、age。英文单词：field，字段。
	    草稿2、行为或者功能：往往是一些动词，即method方法，比如说：liuwan、fei
	*/
	cat := Cat{
		name:"小花",
		age:2,
	}
	fmt.Println(cat.name,cat.age)
	cat.zhualaoshu()

	p := cat//指针赋值
	p.zhuohaozi()
	fmt.Printf("%T\n",p)
	fmt.Println(p)
	fmt.Println(&p)
	//cat.zhuohaozi()
	//以上两种方式都可以实现，指针也以这样做！！！，切记切记！！！！！！！！！！！！
}

type Cat struct {
	name string
	age int
}
//zhuolaoshu这个method接收的类型是一个结构体类型
func (c Cat) zhualaoshu(){
	fmt.Println("猫捉老鼠")
}

//指针: 存储变量的地址。  表示*Cat
func (c *Cat) zhuohaozi(){
	fmt.Println(c.name +" 捉老鼠")
}
