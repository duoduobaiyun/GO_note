package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
  http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
	  writer.Write([]byte("你开心就好"))
  })

  http.HandleFunc("/wuhan/jiayou",jiayou)

  err:=http.ListenAndServe(":8080",nil)//这里的127.0.0.1可以不写，默认值是127.0.0.1，也就是系统的默认地址
	if err!=nil {
		log.Fatal(err)
		return
	}
}

func jiayou(w http.ResponseWriter, v *http.Request)  {
	fmt.Fprintf(w,"武汉加油",2)
}