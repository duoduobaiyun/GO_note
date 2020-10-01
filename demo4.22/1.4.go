
package main

import (
	"fmt"
	"sync"
	"time"
)

var rw0 sync.Mutex
var produce0 int
var wg0 sync.WaitGroup
func main() {
	produce0=50
	//作业一：模拟工厂生产和消费者消费的场景。有一个工厂，生产日用品，每8ms生产一件商品，然后进入库存；
	//另外有三个销售渠道来销售这批产品，三个销售渠道均是每20ms销售一件商品。当库存达到300件时，仓库满了，停止生产程序结束；
	//当库存为0时，没有库存，程序也停止执行。编程模拟实现上述场景。
	//cho:=make(chan bool)
	//cch:=make(chan bool)
	wg0.Add(4)
	go Produce0()
	go Sale0("销售1")
	go Sale0("销售2")
	go Sale0("销售3")
	wg0.Wait()
	fmt.Println("main...over...")
}

func Produce0()  {
	defer wg0.Done()
	for {
		rw0.Lock()
		if produce0 <300 {
			produce0++
			fmt.Printf("生产了第%d件\n",produce0)
			rw0.Unlock()
			time.Sleep(5*time.Millisecond)
		}else if produce0 == 300 {
			rw0.Unlock()
			fmt.Println("库存已满")
			break
		}
	}
}

func Sale0( name string )  {
	defer wg0.Done()
	for {
		rw0.Lock()
		if produce0 >0 {
			produce0--
			fmt.Printf("%s销售第%d件\n",name,produce0)
			rw0.Unlock()
			time.Sleep(20*time.Millisecond)
		}else if produce0 ==0 {
			rw0.Unlock()
			fmt.Println("已售完")
			break
		}
	}
}