package main

import (
	"fmt"
	"sync"
	"time"
)

var x int
var RW sync.RWMutex
var wg sync.WaitGroup
func main() {
	//作业一：
	//有一个工厂，生产日用品，每8ms生产一件商品，然后进入库存；
	//另外有三个销售渠道来销售这批产品，三个销售渠道均是每20ms销售一件商品。
	//当库存达到300件时，仓库满了，停止生产程序结束；
	//当库存为0时，没有库存，程序也停止执行
	wg.Add(3)
	go market()
	go market()
	go market()
	for  {
		time.Sleep(8*time.Millisecond)
		if x<300 {
			x++
			fmt.Printf("生产一件共有%d件\n",x)
		}else if x==300 {
			break
		}
	}
	wg.Wait()


}
func market()  {
	defer wg.Done()
	for  {
		RW.Lock()
		time.Sleep(20*time.Millisecond)
		if x>0 {
			x--
			fmt.Printf("出售一件剩余%d件\n",x)
			RW.Unlock()
		}else if x<=0 {
			RW.Unlock()
			break
		}
	}
}
