package main

import "fmt"

func main() {

func(n int){
        count:=1
		for i:=1;i<=6 ; i++ {
			count*= i
		}
	fmt.Println(count)
	}(6)
}

