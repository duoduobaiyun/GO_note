package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpu:=runtime.NumCPU()
	fmt.Println(cpu)
    runtime.GOMAXPROCS(5)
}