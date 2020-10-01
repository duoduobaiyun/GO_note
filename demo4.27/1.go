package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var c  chan int
var wg sync.WaitGroup
func main() {
	//作业一：启动一个goroutine，生成100个随机数发送到容量为5的缓冲通道ch1中，
	//随机数间隔20ms生成一个；启动另外一个goroutine，从ch1中读取数据，
	//然后计算读取到的数值的3次方，放到通道ch2中；在main函数中，从通道ch2中读取数据，并打印输出。最后程序结束。
	wg.Add(2)
   ch1:=make( chan int ,5)
   ch2:=make( chan int)
   go Num(ch1)
   go Num1(ch1,ch2)
	for a := range ch2{
		fmt.Println(a)
	}
   wg.Wait()

}

func Num(  ch1 chan <- int)  {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<5 ; i++ {
		time.Sleep(20*time.Millisecond)
		fmt.Println(rand.Intn(100))
	}
	close(ch1)
	fmt.Println(ch1)
}

func Num1( ch1  <- chan int,ch2 chan int)  {
	defer wg.Done()
	for{
	    a,ok:= <-ch1
		if !ok {
		  break
		}
		b:=a*a*a
		ch2<-b
	}
	close(ch2)
	fmt.Println(ch2)
}