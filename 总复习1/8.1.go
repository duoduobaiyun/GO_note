package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticketNum int
//read 读   write 写
var rwMutex sync.RWMutex //读写锁
var wg sync.WaitGroup//同步等待组
func main() {
	ticketNum = 500//门票
	/**
	  20000人浏览 --> 1000人下单
	  1、只有1000个人的操作，影响了数据库中的门票信息（修改-1）
	  草稿2、剩余19000个人的操作，没有影响数据库中的门票信息

	  Mutex：互斥锁
	  目的：提高系统的性能（效率）
	*/
	wg.Add(2)

	go DaMai() //一台机器
	go DaMai() //一台机器

	wg.Wait()
	fmt.Println("over ")

}

/**
  Mutex：互斥锁
  修改数据、数据库：写操作
  查询数据、浏览数据：读操作
*/
func DaMai(){

	//妙用defer语句
	defer  wg.Done()

	rand.Seed(time.Now().UnixNano())//随机种子
	for {
		action := rand.Intn(100) // 100以内的数
		//浏览人数：购买人数  =  3 ：1
		if action % 4 == 1 || action % 4 == 2 || action % 4 ==3 {//该人是浏览的

			rwMutex.RLock() //读操作上锁

			time.Sleep(10 * time.Millisecond)
			fmt.Printf("浏览余票信息，剩余%d张\n",ticketNum)

			rwMutex.RUnlock() //读操作解锁

		}else if action % 4 == 0{//该人是购买用户

			rwMutex.Lock() //写上锁

			if ticketNum >0 {//查看余票信息
				time.Sleep(20 * time.Millisecond)
				ticketNum--
				fmt.Printf("大麦网售出一张，剩余%d\n",ticketNum)
			}else if ticketNum <=0{
				rwMutex.Unlock()
				break
			}
			rwMutex.Unlock()
		}
	}
}

