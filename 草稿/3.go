package main

import (
	"fmt"
	"net/http"
)

/**
  自己搭建一个服务器，处理浏览器中login.html的请求
*/
func main() {

	http.HandleFunc("/login",loginFunc1)

	http.ListenAndServe("127.0.0.1:8899",nil)
}

//处理login.html的提交的登录信息
func loginFunc1(writer http.ResponseWriter,request *http.Request){
	//接收数据，解析提交的数据
	request.ParseForm() //解析请求携带的数据 将数据放入到Form当中

	for k,v := range request.Form{
		fmt.Println(k,v)
	}
	fmt.Fprintf(writer,"欢迎你："+request.Form.Get("user"))
}
