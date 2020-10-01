package main

import "fmt"

func main() {
	totallevle := 5
	for i := 1; i <= totallevle; i++ {
		for k := 1; k <= totallevle-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
