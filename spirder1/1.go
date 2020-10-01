package main

import (
	"log"
	"net/http"
)

func main() {

}

func RequestHtml()  {
	request,err:=http.NewRequest("GET","https://movie.douban.com/top250?start=25&filter=",nil)
	if err!=nil {
		log.Fatal(err)
		return
	}
}