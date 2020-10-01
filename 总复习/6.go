package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    ch1:=make(chan int,5)
    ch2:=make(chan int)
	go Num1(ch1)
	go Num2(ch1,ch2)
	for val:= range  ch2{
		fmt.Printf("3次方的值是%d\n",val)
	}
}

func Num1(ch chan int)  {
	defer close(ch)
	rand.Seed(time.Now().UnixNano())
	for i:=1;i<100 ; i++ {
		time.Sleep(20*time.Millisecond)
		a:=rand.Intn(100)
		fmt.Println(a)
	    ch<-a
	}
}

func Num2(ch chan int,ch2 chan int)  {
	defer close(ch2)
	for value:= range ch {
		result:=value*value*value
		ch2<-result
	}
}
