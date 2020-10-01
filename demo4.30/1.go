package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {

}

//fmt.Println("启动监听服务")
func(mux *MyMux) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path:=request.URL.Path
	fmt.Println(path)

	if path == "/userinfo"  {
       Userinfo1(writer,request)
		return
	}else if path == "/student"{
		Student(writer,request)
		return
	}
}
func main() {
//作业一:自定义一个路由处理器，并使用该路由复用器处理http请求，使用自定义的路由处理器处理请求，
//要求web监听的端口是8090。  ①  请求路径：/userinfo   返回内容：查询用户信息 ②  请求路径：/student   返回内容：查询学生信息
    mux:=&MyMux{}
    err:=http.ListenAndServe("127.0.0.1:8090",mux)
	if err!=nil {
		fmt.Println(err.Error())
	}
}

func Userinfo1(writer http.ResponseWriter,request *http.Request)  {
         writer.Write([]byte("查询用户信息"))
}

func Student(writer http.ResponseWriter,request *http.Request)  {
	     writer.Write([]byte("查询学生信息"))
}