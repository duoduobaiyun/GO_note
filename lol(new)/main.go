package main

import (
	"fmt"
	"lol(new)/db_mysql"
	"lol(new)/handfunc"
	"lol(new)/network_request1"
	"net/http"
	"os"
)

func main() {
	fmt.Println("欢迎来到英雄联盟系统")

	//打开数据库
	err := db_mysql.OpenDB() //判断打开数据是否有错误
	if err != nil {
		fmt.Println("打开数据库遇到错误", err.Error())
		os.Exit(0)
	}

	//关闭数据库
	defer db_mysql.CloseDB()

	url := "https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js"
	//查询英雄数量
	heroCount, err := db_mysql.QueryHeroCount()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if heroCount <= 0 {
		//获取并解析网络数据
		herolist, err := network_request1.GetHeroList(url)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//打开事务,并存储每一条数据信息
		tx, err := db_mysql.HeroDB.Begin() //打开事务
		for i := 0; i < len(herolist.Hero); i++ {
			_, err := db_mysql.SaveHeroInfo2DB(herolist.Hero[i])
			if err != nil {
				fmt.Println(err.Error())
				tx.Rollback() //事务回滚
			}
		}
		tx.Commit()
	}
	fmt.Println("英雄的数量",heroCount)

	//6.启动浏览器服务器并在网页上展示
	http.HandleFunc("/login",handfunc.Index)
	http.HandleFunc("/getNewcode",handfunc.Login)

}
