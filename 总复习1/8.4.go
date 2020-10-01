
package main

import (
	"fmt"
	"sync"
	"time"
)
var op sync.WaitGroup
var a int
var b sync.Mutex
func main() {
	a=300
	op.Add(4)
	go production()
	go sell()
	go sell1()
	go sell2()
	op.Wait()
	fmt.Println("最后还剩下：",a)

}
func production()  {
	for   {
		b.Lock()
		if a<301&&a>0 {
			time.Sleep(2*time.Millisecond)
			a++
			fmt.Printf("生产成功，剩余%d张\n",a)
		}else {
			op.Done()
			b.Unlock()
			break
		}
		b.Unlock()
	}

}
func sell()  {
	for {
		b.Lock()
		if a>0 {
			time.Sleep(40*time.Millisecond)
			a--
			fmt.Printf("出售成功，剩余%d张\n",a)
		}else  {
			op.Done()
			b.Unlock()
			break
		}
		b.Unlock()
	}

}
func sell1()  {
	for   {
		b.Lock()
		if a>0 {
			time.Sleep(30*time.Millisecond)
			a--
			fmt.Printf("出售成功，剩余%d张\n",a)
		}else  {
			op.Done()
			b.Unlock()
			break
		}
		b.Unlock()
	}

}
func sell2()  {
	for  {
		b.Lock()
		if a>0 {
			time.Sleep(50*time.Millisecond)
			a--
			fmt.Printf("出售成功，剩余%d张\n",a)
		}else  {
			op.Done()
			b.Unlock()
			break
		}
		b.Unlock()
	}

}