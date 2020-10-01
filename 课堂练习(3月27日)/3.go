package main

import "fmt"

func main() {
	/*
		go语言的数据类型:
		    基本数据类型:
		               int,float,bool,string
		    复合数据类型:
		               array,slice,map,function,pointer,struct,interface。。。
		函数的类型:
		          func(参数列表的数据类型)(返回值列表的数据类型)

	*/
	a:=10
	fmt.Printf("%T\n",a)//int
	b:=[4]int{1,2,3,4}
	fmt.Printf("%T\n",b)//[4]int
	/*
		[4]string
		[6]float64
	*/
	c:=[]int{1,2,3,4}
	fmt.Printf("%T\n",c)

	d:=make(map[int]string)
	fmt.Printf("%T\n",d)
	/*
		map[string]string
		map[string]map[int]string
	*/
	fmt.Printf("%T\n",funa)//func()
	fmt.Printf("%T\n",funb)//func(int) int
	fmt.Printf("%T\n",fund)//func(float64, int, int) (int, int)
	fmt.Printf("%T\n",fune)//func(string, string, int, int) (string, int, float64)
}
func funa()  {}

func funb(a  int)int  {
	return 0
}

func fund(a float64,b,c int)(int,int)  {
	return  0 ,0
}
func fune(a,b string,c,d int)(string,int,float64) {
    return "",0,0
}
