package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//作业一：在一个子协程中打印输出0-1000的随机整数，每次打印输出前，首先睡眠200ms，
	//打印99个数字；在另一个子协程中，每间隔100ms输出一个一个大写字母（范围A-Z），打印199个字幕。
	//在main中等待两个子协程任务执行完毕
	go fun()
	go fun1()

	time.Sleep(2*time.Second)
}

func fun()  {
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<100 ; i++ {
		fmt.Println("——————>",rand.Intn(1000))
		time.Sleep(200)
	}
}

func fun1()  {
	rand.Seed(time.Now().UnixNano())
			for a:=0 ;a<199 ; a++ {
			fmt.Printf("%c\n",rand.Intn(26)+65 )
				time.Sleep(100)
			}
}