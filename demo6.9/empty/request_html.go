package empty

import (
	"io/ioutil"
	"log"
	"net/http"
)

func RequestHtmlPage()(string,error)  {
	client:=http.Client{}
	request,err:=http.NewRequest("GET","https://movie.douban.com/top250?start=0&filter=",nil)
	if err!=nil {
		log.Fatal(err)
		return "",err
	}
	request.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Add("Accept-Language","zh-CN,zh;q=0.9,en;q=0.8,vi;q=0.7")
	request.Header.Add("Cache-Control","no-cache")//缓存的设置：不适用缓存
	request.Header.Add("Connection","keep-alive")//连接：保持连接
	request.Header.Add("Host","movie.douban.com")
	request.Header.Add("Pragma","no-cache")
	request.Header.Add("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36")


	response,err:=client.Do(request)
	if err!=nil {
		log.Fatal(err)
		return "",err
	}
	defer response.Body.Close()

	pageBytes,err:=ioutil.ReadAll(response.Body)
	if err!=nil {
		log.Fatal(err)
		return "",err
	}
	return string(pageBytes),nil
}
