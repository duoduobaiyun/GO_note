package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	/**
	  1、client对象
	  草稿2、request
	  3、client.Do
	  4、处理Response
	*/
	clinet:=http.Client{}
	request,err:=http.NewRequest("get","http://127.0.0.1:9000/login",nil)
	if err!=nil{
        //fmt.Println(err.Error())
		log.Fatal(err)
		return
	}
	reponse,_:=clinet.Do(request)//响应回复请求的内容
	defer reponse.Body.Close()//响应根据请求的内容，从而回复，回复完，然后关闭
	if reponse.StatusCode==200 {
		body,_:=ioutil.ReadAll(reponse.Body)
		fmt.Println(string(body))
	}
}