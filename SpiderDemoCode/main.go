package main

import (
	"SpiderDemoCode/empty"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

var MOVIE_ADD_NUM = 0 //新增电影的数量
var MOVIE_UPDATE_NUM = 0 //更新电影信息的数量
var MOVIE_UNCHANGE_NUM = 0 //电影信息没有发生改变的数量

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

	fmt.Println("开始抓取电影数据...")
	start := time.Now().UnixNano()
	//2、依次爬取第0页，第1页，第2页...
	for pageNum := 0; pageNum < page; pageNum++{
		//请求数据
		htmlString := RequestMovieHtml(pageNum)
		//fmt.Println(htmlString)

		//将网页数据作为数据源，爬取数据, 获取爬取到的数据
		//电影集合：每部电影的信息  结构体
		//结构体格式：{学生编号 学生姓名 学生性别}
		//1291546 霸王别姬  9.6 ...
		//1292052 肖申克的救赎 9.7 ...
		moviesInfo := ParseHtml(htmlString)//htmlString 是一个实参
		for i :=0; i< len(moviesInfo); i++{
			//fmt.Println(moviesInfo[i])
			//moviesInfo[i]代表每一部电影信息，要把moviesInfo[i]存储到数据库中
			//要求和目标：将数据转入到数据库中
			/**
			 * 1、先判断数据库中是否已经存在数据
			 * 2、如果不存在，则调用保存
			 * 3、如果存在，则调用修改
			 */
			exist,_ := empty.QueryMovieInfo(moviesInfo[i].Id)
			if exist{//电影数据存在 true 已存在
				//updateRow变量代表的是sql语句执行后影响的数据库记录的行数
				updateRow,err := empty.UpdateMovieInfo(moviesInfo[i])
				if err != nil {
					log.Fatal(err.Error())
					os.Exit(0)//遇到错误，打印错误，程序停止运行
				}
				if updateRow > 0{//表示数据库中有记录被修改
					MOVIE_UPDATE_NUM++ //更新电影信息的计数+1
					fmt.Println("电影【"+moviesInfo[i].Name+"】更新成功")
				}else{//表示数据库中没有记录被修改
					MOVIE_UNCHANGE_NUM++//未发生修改的数据计数+1
					fmt.Println("电影【"+moviesInfo[i].Name+"】未发生修改")
				}
			}else{//电影数据不存在 false
				id,err :=empty.SaveMovieInfo(moviesInfo[i])
				if err != nil {
					log.Fatal(err.Error())
					os.Exit(0)//如果遇到错误，直接让程序退出
				}
				if id > 0{//每保存成功一次，代表爬取成功一条记录 记录一下
					MOVIE_ADD_NUM++
					fmt.Println("电影 【"+moviesInfo[i].Name+"】保存成功")
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

//单一职责原则：程序功能实现时，遵循模块化，功能核心化的原则。
//现实生活：
//例如：公交车：三个服务人员, 一个驾驶员，一个安全员，一个售票员。
//酒店：后厨，


/**
 * 数据解析函数
 * 功能：以html页面源码为数据源，根据自己编写的爬虫规则，筛选数据，并返回
 */
func ParseHtml(html string) []empty.MovieInfo{
	//ids := make([]string,0)
	//电影编号
	movieIDReg := regexp.MustCompile(`<a href="https://movie.douban.com/subject/(.*?)/" class="">`)
	idSlice := movieIDReg.FindAllStringSubmatch(html,-1)//长度25
	//fmt.Println(idSlice)
	//for _,value := range idSlice  {
	//	ids = append(ids,value[1])
	//}

	//names := make([]string,0)
	//电影名字
	nameReg := regexp.MustCompile(`<img width="100" alt="(.*?)"`)
	nameSlice := nameReg.FindAllStringSubmatch(html,-1)
	//for _,value := range nameSlice  {//长度25
	//	names = append(names,value[1])
	//}

	//rates := make([]string,0)
	//评分
	rateReg := regexp.MustCompile(`<span class="rating_num" property="v:average">(.*?)</span>`)
	rateSlice :=rateReg.FindAllStringSubmatch(html,-1)//-1代表的是筛选数据源中所有的符合规则的数据
	//for _,value := range rateSlice  {//长度25
	//	rates = append(rates,value[1])
	//}

	//for
	movies := make([]empty.MovieInfo,0)
	for i :=0; i<len(idSlice);i++{
		id := idSlice[i][1]//每一部电影的编号
		name := nameSlice[i][1]//每一部电影的名字
		rate := rateSlice[i][1]//每一部电影的评分
		//创建结构体对象
		movie := empty.MovieInfo{Id:id,Name:name,Rate:rate}
		//将结构体对象放入结构体切片中
		movies = append(movies,movie)
	}
	return movies
	//fmt.Println(ids)
	//fmt.Println(names)
	//fmt.Println(rates)
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
	fmt.Println("欢迎进入豆瓣电影爬虫V1.4系统")
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

