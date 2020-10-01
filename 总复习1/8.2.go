package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var tickets int
var wg1 sync.WaitGroup
var rwl sync.RWMutex

func main() {
	tickets = 10000
	wg1.Add(2)
	go dx()
	go dx()
	wg1.Wait()
	fmt.Println("主函数结束")
}

func dx() {
	defer wg1.Done()
	rand.Seed(time.Now().UnixNano())
	for {
		action1 := rand.Intn(1000)
		if action1%5 == 1 || action1%5 == 2 || action1%5 == 3 {
			rwl.RLock()
			time.Sleep(1 * time.Millisecond)
			fmt.Printf("浏览余票信息%d\n", tickets)
			rwl.RUnlock()
		} else if action1%5 == 0||action1%5 == 4 {
			rwl.Lock()
			if tickets > 0 {
				time.Sleep(2 * time.Millisecond)
				tickets--
				fmt.Printf("已售一票,还有%d票\n", tickets)
			} else if tickets <= 0 {
				rwl.Unlock()
				break
			}
			rwl.Unlock()
		}
	}
}
