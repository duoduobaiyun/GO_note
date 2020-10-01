package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)
var BASEURL = "https://movie.douban.com/top250"
var Header = map[string]string{
	"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
	"Accept-Language": "zh-CN,zh;q=0.9",
	"Cache-Control": "max-age=0",
	"Connection": "keep-alive",
	"Host": "movie.douban.com",
	"Pragma": "no-cache",
	"User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.25 Safari/537.36 Core/1.70.3756.400 QQBrowser/10.5.4039.400",
}
defer .Body.Close()

func main() {
	request,err:=http.NewRequest("GET","https://movie.douban.com/top250?start=0&filter=",nil)
	if err!=nil {
		log.Fatal(err)
		return
	}
	client:=http.Client{}
	//添加请求头，模仿浏览器
	//可以接受什么格式的数据

	//可以接受的语言
	request.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language","zh-CN,zh;q=0.9")
    request.Header.Add("Cache-Control","max-age=0")
	request.Header.Add("Connection","keep-alive")
	request.Header.Add("Host","movie.douban.com")
	request.Header.Add("Pragma","no-cache")
	request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.25 Safari/537.36 Core/1.70.3756.400 QQBrowser/10.5.4039.400")

   response,err:=client.Do(request)
	if err!=nil {
		log.Fatal(err)
		return
	}

    htmlBytes,err:=ioutil.ReadAll(response.Body)
	if err!=nil {
		log.Fatal(err)
		return
	}

	html:=string(htmlBytes)

	//编号
	id:=regexp.MustCompile(` <a href="https://movie.douban.com/subject/(.*?)/">`)
	idSlice:=id.FindAllStringSubmatch(html,-1)
    //片名
    name:=regexp.MustCompile(` <img width="100" alt="(.*?)">`)
    nameSlice:=name.FindAllStringSubmatch(html,-1)
    //评分
	ratReg:=regexp.MustCompile(`<span class="rating_num" property="v:average">(.*?)</span>`)
	ratRegSlice:=ratReg.FindAllStringSubmatch(html,-1)
	//评价人数：投票vote
	voteReg:=regexp.MustCompile(`<span>(.*?)人评价</span>`)
	voteRegSlice:=voteReg.FindAllStringSubmatch(html,-1)
	//描述
	desc:=regexp.MustCompile(`<span class="inq">(.*?)</span>`)
	descSlcie:=desc.FindAllStringSubmatch(html,-1)
	//封面
	image:=regexp.MustCompile(`<img src="(.*?)" class="">`)
	imageSlice:=image.FindAllStringSubmatch(html,-1)
	fmt.Println("电影编号   电影名称  评分  评价人数  一句话总结  封面图")

	for i:=0;i<len(nameSlice) ; i++ {
		fmt.Printf("%s %s %s %s %s %s\n",
			idSlice[i][1],
			nameSlice[i][1],
			ratRegSlice[i][1],
			voteRegSlice[i][1],
			descSlcie[i][1],
			imageSlice[i][1])
	}
}
