package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lo(1)/mysql"
	"lo(1)/object"
	"net/http"
	"os"
)

func main() {
	//打开数据库
	err := mysql.OPenDB()
	if err != nil {
		fmt.Println("打开数据库遇到错误", err.Error())
		os.Exit(0)
	}

	//关闭数据库
	defer mysql.CloseDB()

	//查看英雄的数量
	HeroCount, err := mysql.QueryHeroCount()
	if err != nil {
		fmt.Println(err.Error())
	}
	if HeroCount <= 0 {
		url := "https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js"
		client := http.Client{}
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		heroHtml, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		var  herolist object.HeroList
		err=json.Unmarshal(heroHtml,&herolist)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
