package main

import (
	"fmt"
	"strings"
)

func main() {
	str:="WUHANJIAYOUZHONGGUOJIAYOU"
	a:=strings.Contains(str,"ZHONGGUO")
	fmt.Println(a)
	fmt.Println(strings.Index(str,"ZHONGGUO"))
}
