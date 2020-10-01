package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	/**
	  net/http包
	*/
	/**
	  第一个参数addr：address的缩写，监听地址
	  127.0.0.1：特指当前编程所在的电脑本机
	*/

	//提供资源
	//handle:处理。处理资源请求
	http.HandleFunc("/hello_tiger", func(writer http.ResponseWriter, request *http.Request) {
		//返回一句话；天王盖地虎，宝塔镇河妖
		//request：请求
		//response：响应
		//将回复的内容，通过响应发送给请求端（浏览器）
		//write：写
		writer.Write([]byte("天王盖地虎，宝塔镇河妖"))//把数据写入到响应的管道当中
	})


	//程序执行时，该程序会在本地计算机的8080端口，一直监听，等待客户端的连接。
	//err := http.ListenAndServe("127.0.0.1:8080", nil)//死等
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	http.HandleFunc("/wuhan/jiayou",jiayou)

	err := http.ListenAndServe(":8080", nil)//阻塞式
	if err != nil {
		log.Fatal(err.Error())
	}

}

func jiayou(w http.ResponseWriter,r *http.Request){
	//fmt.Println("武汉加油")
	fmt.Fprintf(w,"武汉加油")
}
