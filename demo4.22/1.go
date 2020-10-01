package main

import (
	"fmt"
	"sync"
	"time"
)

var rw sync.Mutex
var wg sync.WaitGroup
var goods int
var ch  chan int
func main() {
	//作业一：模拟工厂生产和消费者消费的场景。有一个工厂，生产日用品，每8ms生产一件商品，然后进入库存；
	//另外有三个销售渠道来销售这批产品，三个销售渠道均是每20ms销售一件商品。当库存达到300件时，仓库满了，停止生产程序结束；
	//当库存为0时，没有库存，程序也停止执行。编程模拟实现上述场景。
	goods = 0
	wg.Add(4)
	go sale("销售1")
	go sale("销售2")
	go sale("销售3")
	go produce()
	r:=<- ch
	if r == 0 {
		r=goods
		fmt.Println("销售完成")
	}
	wg.Wait()
	time.Sleep(50*time.Millisecond)
	fmt.Println("main...over...")
}

func sale(name string) {
	defer wg.Done()
	for {
		rw.Lock()
		time.Sleep(20 * time.Millisecond)
		goods--
		fmt.Printf("%s销售了第%d件\n", name, goods)
		if goods <= 0 {
			ch <- 0
			rw.Unlock()
			//fmt.Println("已售完")
			break
		}
		rw.Unlock()
	}
}

func produce() {
	defer wg.Done()
	for {
		time.Sleep(8 * time.Millisecond)
		goods++
		fmt.Printf("生产第%d个\n", goods)
		if goods == 300 {
			fmt.Println("库存已满")
			break
		}
	}
}
