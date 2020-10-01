package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/userlogin", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("欢迎访问用户登录功能"))
	})

	err := http.ListenAndServe("127.0.0.1:9001", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	a:=http.Client{}
	request,err:=http.NewRequest("get","http://127.0.0.1:9001/userlogin",nil)
	if err!=nil {
		log.Fatal(err)
		return
	}
	reponse ,err := a.Do(request)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer reponse.Body.Close()
	if reponse.StatusCode == 200{
		body,_ := ioutil.ReadAll(reponse.Body)
		fmt.Println(string(body))
	}

}
