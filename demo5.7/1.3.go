package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client:=http.Client{}
	request,_:=http.NewRequest("GET","http://127.0.0.1:9001/userlogin",nil)
	reponse,_:=client.Do(request)
	if reponse.StatusCode == 200 {
	   body,_:=ioutil.ReadAll(reponse.Body)
	   fmt.Println(string(body))
	}
}
