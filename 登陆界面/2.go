package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client:=http.Client{}
	request,err:=http.NewRequest("post","http://127.0.0.1:9001/user_info",nil)
	if err!=nil {
		fmt.Println(err.Error())
	}
	resp,_:=client.Do(request)
	if resp.StatusCode == 200 {
	body,_:=ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	}
}
