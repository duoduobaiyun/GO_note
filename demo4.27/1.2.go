package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var a,b  chan bool
var wg sync.WaitGroup
func main() {
	//作业一：启动一个goroutine，生成100个随机数发送到容量为5的缓冲通道ch1中，
	//随机数间隔20ms生成一个；启动另外一个goroutine，从ch1中读取数据，
	//然后计算读取到的数值的3次方，放到通道ch2中；在main函数中，从通道ch2中读取数据，并打印输出。最后程序结束。
	wg.Add(2)
	ch1:=make( chan int,5)
	ch2:=make( chan int,5)
	go Num()
	go Num1()
	wg.Wait()
	if  {
		
	}
}

func Num()  {
	defer wg.Done()
	var a  int
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<5 ; i++ {
		a= rand.Intn(100)
		time.Sleep(20*time.Millisecond)
	    ch1<-a
	}
}

func Num1()  {
	a:=0
}
	for{
		a=<-ch1
		fmt.Printf("随机数为%d",a)
		ch2 <-a*a*a
		if len(ch1)==0&&d{
			break
		}
}
