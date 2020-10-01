package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("欢迎回来，英雄"))
	})
	err:=http.ListenAndServe("127.0.0.1:9000",nil)
	if err != nil{
		fmt.Println(err.Error())
	}
}