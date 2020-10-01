package main

import (
	"fmt"
	"net/http"
)

func main() {
	//作业一：  ①  编写一个web服务器程序，在自己电脑上的9001端口进行监听，可以处理/userlogin请求,返回一句话“欢迎访问用户登录功能”
	//②  编程实现一个客户端程序，请求步骤①中的服务器，并访问/userlogin，在客户端程序中接收请求，并打印出请求到的信息。
	//程序编写好后，先运行服务器程序，然后运行客户端程序，查看程序运行效果。
	Sever:=http.NewServeMux()
	Sever.HandleFunc("/userlogin", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("欢迎访问用户登录功能"))
	})
	fmt.Println(" 服务启动了，正在监听...")
	http.ListenAndServe("127.0.0.1:9001",Sever)

	//Client:=http.DefaultClient
	//request,_:=http.NewRequest("GET","http://127.0.0.1:9001/userlogin",nil)
	//resp,_:=Client.Do(request)
	//fmt.Println(resp.StatusCode)
	//println(Sever)
	//time.Sleep(20*time.Millisecond)
}

