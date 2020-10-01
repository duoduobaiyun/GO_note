package main

import (
	"fmt"
	"strings"
)

func main() {

	a:="http://192.168.10.100:8080/Day33_Servlet/aa.jpeg"
	b:=strings.Split(a,"/")
	fmt.Println("文件名称:",b[4:5])
	fmt.Println(b)
	fmt.Println(len(b))
	c:=strings.Split(a,".")
    fmt.Println("文件类型:",c[4:5])
}
