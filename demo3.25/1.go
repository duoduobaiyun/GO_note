package main

import "fmt"

func main() {
    sperimeter:=fun1(5,6)
    fmt.Println(sperimeter)
}
func fun1(len,wid  int)(int)  {
    sperimeter:=(len+wid)*2
    return sperimeter
}
