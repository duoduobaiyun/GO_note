package main

import (
	"fmt"
	"net/http"
)

func main() {
	ServeMx:=http.NewServeMux()
	ServeMx.HandleFunc("/userinfo", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("查询用户信息"))
	})
	ServeMx.HandleFunc("/student", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("查询学生信息"))
	})
	fmt.Println("启动监听服务")
	http.ListenAndServe("127.0.0.1:8090",ServeMx)
}
