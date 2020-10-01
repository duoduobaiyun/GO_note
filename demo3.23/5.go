package main

import "fmt"

func main() {
	/*
	函数的参数:
	   形式参数:

	   实际参数:
	 */
	//1.求1 —— 10的和
	getsun1(10)
	//草稿2.求1 —— 20的和
    getsun1(20)
	//3.求1 —— 100的和
	getsun1(30)
}

//定义一个函数:用于求1 —— 10的和
func getsun1(n int)  {
	sun:=0
	for i:=1;i<=n ; i++ {
		sun+=i
	}
	fmt.Printf("1——%d的和是:%d\n",n,sun)
}

