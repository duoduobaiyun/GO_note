package main

import (
	"demo6.12/network"
	"demo6.12/parse"
	"demo6.12/spider_db"
	"fmt"
	"net/http"
	"os"
)

func main() {
	//连接数据库
	err := spider_db.OpenDb()
	if err != nil {
		fmt.Println("打开数据库遇到错误:", err.Error())
		os.Exit(0)
	}

	movieCount, err := spider_db.QueryMovieCount()//数据>0,直接展示即可
	if err != nil {
		fmt.Println(err.Error())
	}
	if movieCount < 0 {//数据库里面没有电影信息,所以要请求数据
		pageHtml, err := network.RequestHtmlPage("https://movie.douban.com/top250")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//2、爬取网页数据
		movies := parse.Parse(pageHtml)

		tx, err := spider_db.MovieDb.Begin() //开启事务
		for i := 0; i < len(movies); i++ {
			//存储每一条数据
			_, err := spider_db.SaveDB(movies[i])
			if err != nil {
				//处理错误
				fmt.Println(err.Error())
				//事务回滚
				tx.Rollback()
			}
		}
		tx.Commit()//提交事务

		//关闭数据库
		err=spider_db.CloseDb()
		if err!=nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Printf("数据库当中有电影数据%d条\n",movieCount)

	//定义浏览器的请求，并指定处理  默认网站的首页往往是：index1.html

	//展示数据：通过浏览器页面
	//web: 开启一个web服务器，接收浏览器的请求，并处理请求
	//addr : address 地址
	fmt.Println("爬虫服务器已在127.0.0.1：9023端口开启，可以接受请求")
	err=http.ListenAndServe(":9023",nil)
	if err!=nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Over")
}