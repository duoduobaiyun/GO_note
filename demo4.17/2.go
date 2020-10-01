package main

import (
	"fmt"
	"sync"
)

var vg sync.WaitGroup


func main() {
    //作业二：在main中启动3个子协程，执行一些耗时操作。
	//使用sync.WaitGroup等待子协程任务完全执行结束后，整个程序执行结束。

	vg.Add(3)

	go fun()
	go fun1()
	go fun2()

	fmt.Println("进入阻塞状态")
	vg.Wait()
	fmt.Println("阻塞状态解除")
}

func fun()  {
	fmt.Println("第一次执行")
	vg.Done()
}

func fun1()  {
	fmt.Println("第二次执行")
    vg.Done()
}

func fun2()  {
	fmt.Println("第三次执行")
	vg.Done()
}