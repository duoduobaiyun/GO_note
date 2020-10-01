package main

import (
	"fmt"
)

func main() {
	/*
	递归函数(recursion):一个函数自己调用自己,就叫做递归函数。
	         递归函数要有一个出口,逐渐的向出口靠近
	 */
	//1.求1-5的和
	sum:=fun1(5)
	fmt.Println(sum)

	//草稿2.fibonacci数列:
	/*
	1  草稿2  3  4  5  6  7    8   9    10   11   12     .....
	1  1  草稿2  3  5  8  13  21   34   55   89   144
	 */
	res:=getfibonacci(6)
	fmt.Println(res)
}

func getfibonacci(n int)int  {
	//fmt.Println("❤")
	if n==1 || n== 2{
		return 1
	}
	return getfibonacci(n-1)+getfibonacci(n-2)
}




func fun1(n int)int  {
	//fmt.Println("❤")
	if n == 1 {
		return 1
	}
	return fun1(n-1)+n
}
/*
求1-5的和
fun1(5)
    fun1(4) + 5
         fun1(3) + 4
              fun1(草稿2) + 3
                   fun1(1) + 草稿2
                    1

 */