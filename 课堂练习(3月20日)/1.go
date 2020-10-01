package main

import "fmt"

func main() {
	a:="坚持坚持坚持再坚持hoqoebqw"
	b:=[]byte(a)
	for _,val:= range  b {
		fmt.Printf("%c\n",val)
	}
}
