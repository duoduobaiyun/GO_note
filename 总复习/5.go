package main

import (
	"fmt"
	"time"
)

func main() {
	for i:=1;i<10 ; i++ {
		fmt.Println(i)
	}

	 go fun()

	time.Sleep(10)
    fmt.Println("main....over....")
}


func fun()  {
	fmt.Println("永别了世界")
}
