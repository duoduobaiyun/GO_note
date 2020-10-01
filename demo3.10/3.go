package main

import (
	"fmt"
)

func main() {
	var sum float64 = 0
   var arr1 =[...]float64{1.2315,2.1566,2.15615,3.1615}
	for index,value :=range arr1{
      fmt.Printf("下标是:%d,数值是:%f\n",index,value)
	}
	//float64(sum)
	for _,v:=range arr1 {
		sum += v
	}
	fmt.Println(sum)
}
