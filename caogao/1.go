package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go Rutine(ch)
	time.Sleep(500*time.Millisecond)
	res := <-ch
	fmt.Println(res)
}
func Rutine(ch1 chan int)  {
	ch1 <- 10
}
