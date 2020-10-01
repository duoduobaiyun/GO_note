package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)
//全局变量
var BASEURL = "https://movie.douban.com/top250"
var Header = map[string]string{
	"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"Accept-Language":"zh-CN,zh;q=0.9,en;q=0.8,vi;q=0.7",
	"Cache-Control":"no-cache",
	"Connection":"keep-alive",
	"Host":"movie.douban.com",
	"Pragma":"no-cache",
	"User-Agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36",
}

func main() {
	/**
	 * ① 输出提示信息，并读取用户的输入：page页面  采集几页
	 * ② 依次采集第1页，第2页，第3页...执行网络请求：豆瓣url，获取到html页面
	 * ③ 解析html页面：爬取数据
	 * ④ 数据存储：把爬取到的数据存储到mysql数据库当中
	 * ⑤ 输出提示信息：爬取的结果（成功/失败)、耗时说明...

	 * fmt:标准的输入输出
	 * 网络编程：http
	 */
	//1、输出帮助和提示信息print
	page, err := PrintHelperInfo()
	if err != nil{
		log.Fatal(err.Error())
		fmt.Println("请重新尝试!")
		os.Exit(0)
	}

	//2、依次爬取第0页，第1页，第2页...
	for pageNum := 0; pageNum < page; pageNum++{
		//请求数据
		htmlString := RequestMovieHtml(pageNum)
		fmt.Println(htmlString)
	}

}

func RequestMovieHtml(page_num int) string{
	client := http.Client{}//客户端
	//实例化request
	offsetNum := strconv.Itoa(page_num * 25)
	request,err := http.NewRequest("GET","https://movie.douban.com/top250?start="+offsetNum+"&filter=",nil)
	if err != nil{
		log.Fatal(err.Error())
		return ""
	}
	//请求：直接请求是请求不到的？网站做了反爬措施
	//如何解决反爬的问题：模拟一个真实的浏览器请求
	request.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Add("Accept-Language","zh-CN,zh;q=0.9,en;q=0.8,vi;q=0.7")
	request.Header.Add("Cache-Control","no-cache")//缓存的设置：不适用缓存
	request.Header.Add("Connection","keep-alive")//连接：保持连接
	request.Header.Add("Host","movie.douban.com")
	request.Header.Add("Pragma","no-cache")
	request.Header.Add("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36")

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	defer resp.Body.Close()
	bytes, err :=ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	return string(bytes)
}

//提示帮助信息
func PrintHelperInfo() (int,error){
	fmt.Println("欢迎进入豆瓣电影爬虫V1.1系统")
	var page int
	fmt.Println("请输入你要采集的页码数字（比如2代表采集前2页,请保持在10以内)：")
	_,err :=fmt.Scan(&page)
	if err != nil{
		return 0,err
	}
	if page > 10{
		return 0,errors.New("只能采集前10页数据，我尽力了")
	}
	return page,nil
}

/**
 * 请求页面信息
 */
func RequestHtmlPage(url string) (string,error){
	client := http.Client{}

	request,err := http.NewRequest("GET",url,nil)
	if err != nil {
		log.Fatal(err.Error())
		return "",err
	}

	//添加header
	for key,value := range Header{
		request.Header.Add(key,value)
	}

	//执行请求
	response,err := client.Do(request)
	if err != nil {
		log.Fatal(err.Error())
		return "",err
	}

	htmlBytes,err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err.Error())
		return "",err
	}
	//通过网络请求把网页数据转换成html的string格式，然后输出
	return string(htmlBytes),nil
}


func ParseMovieInfo(html string){

}
