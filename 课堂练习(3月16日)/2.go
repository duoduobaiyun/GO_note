package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/**
	  定义一个切片,对切片添加5个10以内的随机值
	*/
	slice := make([]int,3)
	fmt.Println(len(slice),cap(slice))

	rand.Seed(time.Now().UnixNano())
	slice = append(slice,rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10))
	fmt.Println(slice)

}
