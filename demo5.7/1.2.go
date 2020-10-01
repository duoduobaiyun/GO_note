//package main
//
//import (
//	"fmt"
//	"net/http"
//)
//
//func main() {
//	http.HandleFunc("/userlogin", func(writer http.ResponseWriter, request *http.Request) {
//		url := request.URL
//		fmt.Println("host:", url.Host)
//		fmt.Println("path", url.Path)
//		fmt.Println("协议类型：", url.Scheme)
//		fmt.Println("请求方法：", request.Method)
//		writer.Write([]byte("欢迎访问用户登录功能"))
//	})
//
//	fmt.Println("开始监听。。")
//	err := http.ListenAndServe("127.0.0.1:9001", nil)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//
//}
