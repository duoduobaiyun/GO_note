package main

import "fmt"

func main() {
	a := []byte{97,98,99,100}
	fmt.Println(a)
	for i:=0;i<len(a) ;i++  {
		//fmt.Printf("%c\n",a)
	}
	fmt.Printf("%c\n",a)
}