package main

import (
	"fmt"
	"net/http"
)

func main() {
	//作业一：编程实现一个web服务器，可以处理名为/login的请求，模拟登陆请求，携带用户名和密码两个数据。在服务器端接收登陆的数据，并对请求数据进行判断。
	//如果用户名是hello，密码是123456，表示用户存在，返回“恭喜登录成功“结果字样，否则返回”对不起，您的账号或者密码不正确，请重新尝试“的提示语。
	fmt.Println(" 启动监听服务 ")
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {

		url := request.URL
		fmt.Println("host:",url.Host)
		fmt.Println("path",url.Path)
		fmt.Println("协议scheme：",url.Scheme)//协议
		fmt.Println("请求方法(类型): ",request.Method)

		username:=request.FormValue("username")
		fmt.Println("账号:",username)
        password:=request.FormValue("password")
        fmt.Println("密码:",password)
		if username == "hello"&& password=="123456"{
			//fmt.Println("恭喜登录成功")
			writer.Write([]byte("恭喜登录成功"))
		}else{
			//fmt.Println("对不起，您的账号或者密码不正确，请重新尝试")
			writer.Write([]byte("对不起，您的账号或者密码不正确，请重新尝试"))
		}
		//writer.Write([]byte("天王盖地虎，宝塔镇河妖"))
	})
	err:=http.ListenAndServe("127.0.0.1:8080",nil)
	fmt.Println(" 服务启动了，正在监听...")
	if err != nil {
		fmt.Println(err.Error())
	}

}
