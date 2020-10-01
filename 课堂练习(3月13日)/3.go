package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	arr:=[10]int{}
	count:=0//计时器
	rand.Seed(time.Now().UnixNano())//种子数的值改变也会随机改里面的数值
	for i := 0; i<10 ; i++ {//循环10次
		 //永真
			lb: //跳到这
			count++//每次加上1
			n:=rand.Intn(15)//随机数[0,15)
			for j:=0;j<=i ; j++ {//内部循环
				if arr[j] == n {//arr[J]是下标 全等于 n是相对应的随机数，也就是说arr[j]的值来判断
					goto lb//跳出
				}
			}
			arr[i] = n//arr[10]的次数=随机数
			//break//跳出循环//原来要把break去掉
		   }
		fmt.Println(arr)
	    fmt.Println(count)
	}
//arr[i]的值是下标，arr[j]是用来判断下标的