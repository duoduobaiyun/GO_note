package main

import (
	"fmt"
	"net/http"
)

func main() {
	//作业二：编写一个web服务，在浏览器中访问该服务时，显示自己硬盘上某个存在的文目录。服务的端口是9000。
	err:=http.ListenAndServe("127.0.0.1:9000",http.FileServer(http.Dir("C:/Users/GOProjects/src/demo4.29")))
	if err != nil {
		fmt.Println(err.Error())
	}
}
