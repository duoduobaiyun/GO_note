package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	uslStr:="https://ke.youdao.com/user/mycourse/"
	response,err:=http.Get(uslStr)
	if err!=nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	statusCode:=response.StatusCode
	if statusCode == 200 {
		body,err:=ioutil.ReadAll(response.Body)
		if err!=nil {
			log.Fatal(err)
			return
		}
		fmt.Println(string(body))

		fmt.Println("header:")
		head:=response.Header
		for  key,val:= range head  {
			fmt.Println(key,val)
		}
		request:=response.Request
		fmt.Println("method:",request.Method)
		fmt.Println("url:",request.URL)
		fmt.Println("path:",request.URL.Path)
		fmt.Println("cookies",request.Cookies())
	}else {
		fmt.Println("状态码错误")
	}
}
