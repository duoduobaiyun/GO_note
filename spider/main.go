package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"spider/out"
	"strconv"
	"time"
)
var MOVIE_ADD_NUM = 0 //新增电影的数量
var MOVIE_UPDATE_NUM = 0 //更新电影信息的数量
var MOVIE_UNCHANGE_NUM = 0 //电影信息没有发生改变的数量
func main() {
	page,err:=PrintHelpInfo()
	if err!=nil {
		log.Fatal(err)
		fmt.Println("请重新尝试")
		os.Exit(0)
	}
	fmt.Println("开始抓取页面")
	start:=time.Now().UnixNano()
	//2、依次爬取第0页，第1页，第2页...
	for pageNum:=0;pageNum<page ; pageNum++ {
		//请求数据
		htmlString:=RequestMovieHtml(pageNum)
		//fmt.Println(htmlString)

		//将网页数据作为数据源，爬取数据, 获取爬取到的数据
		//电影集合：每部电影的信息  结构体
		//结构体格式：{学生编号 学生姓名 学生性别}
		//1291546 霸王别姬  9.6 ...
		//1292052 肖申克的救赎 9.7 ...
		movieInfo:=ParseHtml(htmlString)
		for i:=0;i<len(movieInfo) ; i++ {
			//fmt.Println(moviesInfo[i])
			//moviesInfo[i]代表每一部电影信息，要把moviesInfo[i]存储到数据库中
			//要求和目标：将数据转入到数据库中
			/**
			 * 1、先判断数据库中是否已经存在数据
			 * 2、如果不存在，则调用保存
			 * 3、如果存在，则调用修改
			 */
			exit,_:=out.QueryMovieInfo(movieInfo[i].Id)
			if exit {//电影数据存在 true 已存在
				//updateRow变量代表的是sql语句执行后影响的数据库记录的行数
				updateRow,err := out.UpdateMovieInfo(movieInfo[i])
				if err!=nil {
					log.Fatal(err)
					os.Exit(0)
				}
				if updateRow>0 {
					MOVIE_UPDATE_NUM++ //更新电影信息的计数+1
					fmt.Println("电影【"+movieInfo[i].Name+"】更新成功")
				}else{//表示数据库中没有记录被修改
					MOVIE_UNCHANGE_NUM++//未发生修改的数据计数+1
					fmt.Println("电影【"+movieInfo[i].Name+"】未发生修改")
				}
			}else {
				//电影数据不存在 false
				id,err :=out.SaveMovieInfo(movieInfo[i])
				if err != nil {
					log.Fatal(err.Error())
					os.Exit(0)//如果遇到错误，直接让程序退出
				}
				if id>0 {//每保存成功一次，代表爬取成功一条记录 记录一下
					MOVIE_ADD_NUM++
					fmt.Println("电影 【"+movieInfo[i].Name+"】保存成功")
				}
			}
		}
	}
	//最后，输出统计信息
	end := time.Now().UnixNano()
	fmt.Println()//空格
	fmt.Println("==================统计信息=====================")
	fmt.Println("豆瓣电影TOP250数据采集结束:")
	fmt.Printf("新增电影数据%d部\n",MOVIE_ADD_NUM)
	fmt.Printf("更新电影数据%d部\n",MOVIE_UPDATE_NUM)
	fmt.Printf("数据未变化电影%d部\n",MOVIE_UNCHANGE_NUM)
	fmt.Printf("本次爬取耗时：%fms \n",float32(end-start)/1000000)
}

func ParseHtml(html string)[]out.MovieInfo  {
	//编号
	movieId:=regexp.MustCompile(`<a href="https://movie.douban.com/subject/(.*?)/">`)
	movieIdSlice:=movieId.FindAllStringSubmatch(html,-1)
	//fmt.Println(len(movieIdSlice))
	//片名
	movieName:=regexp.MustCompile(` <img width="100" alt="(.*?)" `)
	movieNameSlice:=movieName.FindAllStringSubmatch(html,-1)
	//fmt.Println(len(movieNameSlice))
	//评分
	rateReg:=regexp.MustCompile(`<span class="rating_num" property="v:average">(.*?)</span>`)
	rateRegSlice:=rateReg.FindAllStringSubmatch(html,-1)
	//fmt.Println(len(rateRegSlice))
	//评分人数
	vote:=regexp.MustCompile(`<span>(.*?)人评价</span>`)
	voteSlice:=vote.FindAllStringSubmatch(html,-1)
	//fmt.Println(len(voteSlice))
	//描述
	desc:=regexp.MustCompile(`<span class="inq">(.*?)</span>`)
	descSlice:=desc.FindAllStringSubmatch(html,-1)
	//fmt.Println(len(descSlice))
	//图片
	image:=regexp.MustCompile(` src="(.*?)" class="">`)
	imageSlice:=image.FindAllStringSubmatch(html,-1)
    //fmt.Println(imageSlice)

	movies:=make([]out.MovieInfo,0)
	for i:=0; i<len(movieIdSlice); i++ {
		id:=movieIdSlice[i][1]
		name:=movieNameSlice[i][1]
		rat:=rateRegSlice[i][1]
		vote:=voteSlice[i][1]
		desc:=descSlice[i][1]
		image:=imageSlice[i][1]
		//创建结构体对象
		movie:=out.MovieInfo{Id:id,Name:name,RatReg:rat,Vote:vote,Desc:desc,Image:image}
		//将结构体对象放入结构体切片中
		movies = append(movies,movie)
	}
    return movies
}

func RequestMovieHtml(page_num int)string  {
	client:=http.Client{}
	//实例化request
	offsetNum:=strconv.Itoa(page_num * 25)
	request,err:=http.NewRequest("GET","https://movie.douban.com/top250?start="+offsetNum+"&filter=",nil)
	if err!=nil {
		log.Fatal(err)
		return ""
	}
	//请求：直接请求是请求不到的？网站做了反爬措施
	//如何解决反爬的问题：模拟一个真实的浏览器请求
	request.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language","zh-CN,zh;q=0.9")
	request.Header.Add("Cache-Control","must-revalidate, no-cache, private")
	request.Header.Add("Connection","keep-alive")
	request.Header.Add("Host","movie.douban.com")
	request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.25 Safari/537.36 Core/1.70.3756.400 QQBrowser/10.5.4039.400")
    request.Header.Add("Pragma","no-cache")

	response,err:=client.Do(request)
	if err!=nil {
		log.Fatal(err)
		return ""
	}
	defer response.Body.Close()

	bytes,err:=ioutil.ReadAll(response.Body)
	if err!=nil {
		log.Fatal(err)
		return ""
	}
	return string(bytes)
}

func PrintHelpInfo()(int,error)  {
	fmt.Println("欢迎来到豆瓣爬虫系统")
	var page int
	fmt.Println("请输入你要采集的页码数字（比如2代表采集前2页,请保持在10以内)：")
	_,err:=fmt.Scan(&page)
	if err!=nil {
		return 0,err
	}
	if page > 10 {
		return 0,errors.New("超出了我的能力范畴")
	}
	return page,nil
}