package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
  //1、相当于我们的浏览器客户端
  client:=http.Client{}

  //2、访问百度，生成一个访问百度的请求  request(请求百度服务器获取数据)
  request,err:=http.NewRequest("GET","http://www.baidu.com",nil)
	if err!=nil {
		log.Fatal(err)
		return
	}
  response,err:=client.Do(request)
	if err!=nil {
		log.Fatal(err)
		return
	}
	defer response.Body.Close()


    //处理返回 状态码正常就执行出来
	if response.StatusCode==200 {
		body,_:=ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
}
