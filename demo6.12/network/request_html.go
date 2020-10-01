package network

import (
	"io/ioutil"
	"log"
	"net/http"
)

func RequestHtmlPage(page string)(string,error)  {
	client:=http.Client{}


	//request
	request,err:=http.NewRequest("GET",page,nil)
	if err!=nil {
		log.Fatal(err)
		return "",err
	}
	request.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language","zh-CN,zh;q=0.9")
	request.Header.Add("Cache-Control","max-age=0")
	request.Header.Add("Connection","keep-alive")
	request.Header.Add("Host","movie.douban.com")
	request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.25 Safari/537.36 Core/1.70.3756.400 QQBrowser/10.5.4039.400")
    request.Header.Add("Pragma","no-cache")

	response,err:=client.Do(request)
	if err!=nil {
		log.Fatal(err)
		return "",err
	}
	defer response.Body.Close()

	pageHtml,err:=ioutil.ReadAll(response.Body)
	if err!=nil {
		log.Fatal(err)
		return "",err
	}
	return string(pageHtml),nil
}
