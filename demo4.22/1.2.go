package main

import (
	"fmt"
	"sync"
	"time"
)

var produce1 int = 100
//推荐使用通道解决协程间的通信问题
//var ch chan bool//通道变量
var mutex1 sync.Mutex
func main() {
	ch := make(chan bool)
	ch1 := make(chan bool)
	go Produce1(ch) //生产
	go Sale1(ch1) //销售
	//读数据
	<-ch //从通道里面读数据 ：一直等待直到有数据被读取出来：阻塞。
	<-ch1 //从通道里面读数据：标准输入输出：rand.Scan 阻塞式。
	fmt.Println("over 。。。结束")
}

func Produce1(ch chan bool) {
	defer  close(ch)//关闭通道
	for {
		mutex1.Lock() //先上锁
		if produce1 < 300 {
			//time.Sleep(1 *time.Millisecond)//生产过程 8毫秒 销售是无法进行的
			produce1++
			fmt.Println("生成了商品，当前库存:", produce1)
			mutex1.Unlock() //释放了锁
			time.Sleep(8 * time.Millisecond)
		} else if produce1 >= 300 {
			fmt.Println("库存满了，停止生产。。。")
			mutex1.Unlock()
			ch <- true //把数据写入到通道中
			break
		}
	}
}

func Sale1(ch chan bool) {
	defer  close(ch)//关闭通道
	for {
		mutex1.Lock()
		if produce1 > 0 {
			//time.Sleep(10 * time.Millisecond) //销售过程
			produce1--
			fmt.Println("销售了商品，当前库存:", produce1)
			mutex1.Unlock()
			time.Sleep(5 * time.Millisecond)
		} else if produce1 <= 0 {
			fmt.Println("没库存了, 停止销售")
			mutex1.Unlock()
			ch <- true //不销售了，没库存了
			break
		}
	}
}

