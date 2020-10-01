package main

import (
	"DoubanNewMovie/network"
	"DoubanNewMovie/parse"
	"DoubanNewMovie/reqhandle"
	"DoubanNewMovie/spider_db"
	"fmt"
	"net/http"
	"os"
)

/**
 * 程序的主入口
 * main函数的要求：main函数所在的软件包名必须是main
 */
func main() {

	//连接数据库操作：打开数据库
	err :=spider_db.OpenDb()
	if err !=nil {
		fmt.Println("打开数据库遇到错误：",err.Error())
		os.Exit(0)
	}

	//一、查询一下数据库中是否已经存在数据了，
	//二、如果数据不存在，先请求网络数据，并保存到数据
	//三、如果数据已经存在，直接展示数据

	//查询数据库当中电影记录的数据量，如果数据量>0，表示有数据，直接展示即可
	movieCount,err :=spider_db.QueryMovieCount()
	if err != nil {
		fmt.Println(err.Error())
	}
	if movieCount <=0 {//如果movieCount小于等于0，表示数据库当中没有电影数据，应该先请求网络并采集数据，存到数据库中
		//1、请求网页数据
		pageHtml,err :=network.ReqeustHtmlPage("https://movie.douban.com/chart")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//2、爬取网页数据
		movies :=parse.ParsePage(pageHtml)

		//3、存储爬取后的数据
		// 一共有10条数据： 前5条没有问题，存第6条数据的时候出现的了问题。
		//要求：当出现操作多条数据的时候，把多条数据看做一个整体。
		//数据回滚。
		//原子：最小单位
		//原子性操作：多条数据操作时，遵循原子性操作，要么全部成功，要么全么失败，执行过程中遇到失败，需要回滚。
		//数据库事务：

		//开启事务
		tx,err :=spider_db.MovieSpiderDb.Begin() //开启事务
		for i := 0; i< len(movies);i++{
			//存储每一条电影记录
			_, err :=spider_db.SavieMovie2Db(movies[i])
			if err != nil {
				//处理错误
				fmt.Println(err.Error())
				//事务回滚
				tx.Rollback()//回滚
			}
		}
		//提交事务
		tx.Commit()//提交事务

		//关闭数据库
		err =spider_db.CloseDB()
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Printf("数据库当中有电影数据%d条\n",movieCount)


	//定义浏览器的请求，并指定处理
	http.HandleFunc("/",reqhandle.Index)

	//展示数据：通过浏览器页面
	//web: 开启一个web服务器，接收浏览器的请求，并处理请求
	//addr : address 地址
	fmt.Println("爬虫服务器已在127.0.0.1:9090端口开启，可以接收请求")
	err =http.ListenAndServe(":9090",nil)//阻塞状态
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Over")
}
