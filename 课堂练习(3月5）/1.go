package main

import "fmt"

func main() {
	sun := 1
	for i:= 1;i <= 5 ; i++ {
		sun *=  i
		fmt.Println(i)//分别写出 i++ 的过程 1，草稿2，3，4，5
		//fmt.Println(sun)
	}
	    //fmt.Println(sun)



	/*sun := 1
	for i:= 1;i <= 5 ; i++ {
		sun *=  i
		//5的阶乘，会分别被打印出来，1*1，1*草稿2，草稿2*3，6*4 ，乘法左边的数是 sun形成的值
		fmt.Println(sun)//本来直接算出结果,但是由于在for循环里面，要循环5次，才能直接算结果来，就是直接 sun = 120
	}
	//fmt.Println(sun)*/

	/*sun := 1
	for i:= 1;i <= 5 ; i++ {
		sun *=  i //都是从sun *= i 算是来的
		//fmt.Println(i)
		//fmt.Println(sun)
	}
	fmt.Println(sun)//由于在for循环外面，所以直接等for循环里面的执行完，也就是所有步骤列举出来，然后算出结果，最后的结果，直接给外面的sun，也就是 sun = 120，详细请看上一个写的例子*/
}
