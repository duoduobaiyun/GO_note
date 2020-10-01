package main

import (
	"fmt"
	"strings"
)

func main() {
	coun:="cwuibiwbebcuobwucbowe"
	funw(coun)
	fmt.Println("卧槽",coun)
}
func funw(coun int)map[string]int  {
	count  :=  make(map[string]int)
	str2:=strings.Split(coun,"")
	for _,v:=range str2 {
		count[v]++
		fmt.Println(count)
	}
	return count
}
