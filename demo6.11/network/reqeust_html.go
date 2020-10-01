package network//程序的包名

import (
	"io/ioutil"
	"net/http"
)

/**
 * 核心功能：发起网络请求，并得到网页数据
 */
func ReqeustHtmlPage(page_url string)(string,error){

	client := http.Client{}//模拟的是浏览器工具

	//request
	request,err := http.NewRequest("GET",page_url,nil)
	if err != nil{
		return "",err
	}

	//网站有反爬措施：将程序的request请求，伪装成为一个真正的浏览器请求
	//Encoding：编码
	request.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Add("Accept-Language","zh-CN,zh;q=0.9,en;q=0.8,vi;q=0.7")
	request.Header.Add("Cache-Control","no-cache")//缓存的设置：不适用缓存
	request.Header.Add("Connection","keep-alive")//连接：保持连接
	request.Header.Add("Host","movie.douban.com")
	request.Header.Add("Pragma","no-cache")
	request.Header.Add("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36")

	//发起请求
	resp,err :=client.Do(request)
	if err != nil{
		return "",err
	}
	defer resp.Body.Close()

	pageBytes,err :=ioutil.ReadAll(resp.Body)
	if err != nil{
		return "",err
	}
	return string(pageBytes),nil
}