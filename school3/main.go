package main

import (
	"fmt"
	"net/http"
	"os"
	"school3/network"
	"school3/parse"
	"school3/reqhandle"
	"school3/spider_db"
)

func main() {
	err :=spider_db.OpenDb()
	if err !=nil {
		fmt.Println("打开数据库遇到错误：",err.Error())
		os.Exit(0)
	}


	movieCount,err := spider_db.QueryMovieCount()
	if err!= nil{
		fmt.Println(err.Error())
	}
	if movieCount<=0 {
		pageHtml , err := network.RequestHtmlPage("https://movie.douban.com/chart")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		movies := parse.ParsePage(pageHtml)
		tx,err :=spider_db.MovieSpiderDb.Begin()
		for i := 0; i< len(movies);i++{

			_, err :=spider_db.SavidMovie2Db(movies[i])
			if err != nil {
				fmt.Println(err.Error())

				tx.Rollback()
			}
		}
		tx.Commit()

	}

	fmt.Printf("获取到%d条的电影信息\n",movieCount)
	http.HandleFunc("/",reqhandle.Index)
	http.HandleFunc("/login",reqhandle.Login)

	fmt.Println("服务器已开启，可以接受请求")
	err = http.ListenAndServe(":9091", nil)
	if err!= nil {
		fmt.Println(err.Error())

	}
	spider_db.CloseDB()//defer
}