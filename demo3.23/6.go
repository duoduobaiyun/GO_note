package main

import "fmt"

func main() {
	/*
		函数的参数:
		   形式参数:

		   实际参数:
	*/
	//1.求1 —— 10的和
	getsun2(10)
	//草稿2.求1 —— 20的和
	getsun2(20)
	//3.求1 —— 100的和
	getsun2(100)
	//3.求2个整数的和
	getsunAdd(10,20)
	getsunAdd2(1,2)

	fun1(1,3,"hello")
}

//定义一个函数:用于求1 —— 10的和
func getsun2(n int)  {
	sun:=0
	for i:=1;i<=n ; i++ {
		sun+=i
	}
	fmt.Printf("1——%d的和是:%d\n",n,sun)
}

func getsunAdd(a int,b int)  {
     sun:=a + b
     fmt.Printf("%d + %d = %d\n",a,b,sun)
}

func getsunAdd2(a,b int)  {//参数的类型一致,可以简写在一起
	fmt.Printf("a:%d,b:%d\n",a,b)
}

func fun1(a,b float64,c string)  {
	fmt.Printf("a:%.2f,b:%.2f,c:%s\n",a,b,c)
}