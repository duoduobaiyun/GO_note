package main

import "fmt"

func main() {
	var height int
	var weight int
	fmt.Scanln(&height,&weight)
	fmt.Printf("身高是:%d cm,体重是:%d 斤",height,weight)
}
