package main

import "fmt"

func main() {

	fun0()

}

func fun0()  {
	for i:=0;i<26 ; i++ {
		fmt.Printf("%c\n",'A'+i)
	}
}