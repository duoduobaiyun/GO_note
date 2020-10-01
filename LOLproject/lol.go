package main

import (
	"LOLproject/mysql_db"
	"LOLproject/parse"
	"LOLproject/reqhandle"
	"LOLproject/request"
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := mysql_db.Open_db()
	if err != nil {
		fmt.Println("打开数据库遇到错误：", err.Error())
		os.Exit(0)
	}
	hereCount, err := mysql_db.SelectRows()
	if err != nil {
		fmt.Println(err.Error())
	}
	if hereCount <= 0 {
		pagehtml, err := request.Request_html("https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		hreoes := parse.Hero_html(pagehtml)
		//开启事务
		affair, err := mysql_db.LOL.Begin()
		for i := 0; i < len(hreoes); i++ {
			_, err := mysql_db.Hero_db(hreoes[i])
			if err != nil {
				fmt.Println(err.Error())
				//事务回滚
				affair.Rollback()
			}
		}
		//提交事务
		affair.Commit()

	}
	fmt.Printf("数据库中有 %d 条英雄数据\n",hereCount)
	http.HandleFunc("/", reqhandle.Log)
	http.HandleFunc("/login", reqhandle.Log1)
	fmt.Println("127.0.0.1:1010端口开启，可以接收请求")
	err = http.ListenAndServe(":1010", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = mysql_db.Close_db()
	if err != nil {
		fmt.Println(err.Error())
	}
}
