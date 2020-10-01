package main

import "fmt"

func main() {
	  a := 0
	  b := 1
	  c := 2
	  a = b
	  b = c
      c = a
	  fmt.Println(b)
}