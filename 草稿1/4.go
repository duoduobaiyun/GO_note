package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client:=http.Client{}
	request,err:=http.NewRequest("GET","https://home.meishichina.com/recipe.html",nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	response,err:=client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	html,err:=ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(html))
}
