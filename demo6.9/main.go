package main

import (
	"demo6.9/empty"
	"fmt"
	"log"
)

func main() {
     pageHtml,err:=empty.RequestHtmlPage()
	if err!=nil {
		log.Fatal(err)
		return
	}
	fmt.Println(pageHtml)
}
