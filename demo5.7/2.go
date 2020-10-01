package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	Client:=http.DefaultClient
	request,_:=http.NewRequest("GET","http://127.0.0.1:9001/userlogin",nil)
	resp,_:=Client.Do(request)
	fmt.Println(resp.StatusCode)
	if resp.StatusCode  == 200{
		body,_:=ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
	//println(Sever)
}
