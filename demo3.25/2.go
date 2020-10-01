package main

import (
	"fmt"
	"strings"
)

func main() {
    fmt.Print(abc("vwvveqvoinowinvownvwo"))
}

func abc(str string)string  {
	//a:="zhongguojiayouwuhangjiayou"
	for i:='a';i<='z' ; i++ {
		a:=string(i)
		n:=strings.Count(str,a)
		if n!=0{
			fmt.Println("这个字符", a ,"一共出现了", n ,"次")
		}
	}
 return str
}