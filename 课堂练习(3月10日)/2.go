package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num := rand.Int()
	fmt.Print(num)

	//for i := 0; i<10 ; i++ {
	//	num := rand.Intn(10)//[0,9]
	//	fmt.Println(num)//打印随机数10次,执行几次都是一样的
	//}
	//    rand.Seed(1)//种子数
	//    num2 := rand.Intn(100)//限制数字的大小，[0,99]
	//    fmt.Println("——>",num2)//结果
	//
	//    t1:=time.Now()
	//    fmt.Println(t1)
	//    fmt.Printf("%T\n",t1)
	//    //时间戳，指定时间，距离1970年1月1日0点0分0秒，之间的时间差值:秒，纳秒
	//    timeStamp1:=t1.Unix()//秒  距离1970年1月1日0点0分0秒到现在
	//    fmt.Println(timeStamp1)
	//
	//    timeStamp2:=t1.UnixNano()//纳秒  距离1970年1月1日0点0分0秒到现在
	//    fmt.Println(timeStamp2)
	//
	   //设置种子数，可以设置成时间戳
	   rand.Seed(time.Now().UnixNano())
	   for i:= 0 ; i < 10 ; i++ {
		//调用生成随机数的函数
	   	fmt.Println("————>",rand.Intn(100))
	}
	///*[15,76]
	//  [3,48]
	//  [0,45]+3
	//
	//  Intn(n) // [0,n]
	// */
	//
	//for i := 0;i<10 ; i++ {
	//	num3:=rand.Intn(46)+3//[3,48]
	//	fmt.Println(num3)
	//	fmt.Println("——————>",num3)
	//	num4:=rand.Intn(62)+15//[15,76]
	//	fmt.Println(num4)
	//}

}
