package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"project1/lo/db_mysql"
	"project1/lo/object"
	"project1/lo/parse"
	"project1/lo/reqhandlet"
)


func main() {
	err:=db_mysql.OPenDB()
	if err!=nil{
		fmt.Println("打开数据库遇到错误：", err.Error())
		os.Exit(0)
	}


    defer 	 db_mysql.CloseDB()

	HeroCount,err:=db_mysql.QueryHeroCount()
	if err != nil {
		fmt.Println(err.Error())
	}
	if HeroCount<=0 {
		url := "https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js"

		client := http.Client{}

		request, err := http.NewRequest("GET", url, nil)//获取
		if err != nil {
			fmt.Println(request)
			return
		}
		resp, err := client.Do(request)
		if err != nil {
			fmt.Println(request)
			return
		}
		heroData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(request)
			return
		}
		//fmt.Println(string(heroData))

		var herolist object.HeroList
		err = json.Unmarshal(heroData, &herolist)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		heros := parse.PageHtml(herolist.Hero)
		//fmt.Println(heros)
		affair, err := db_mysql.LolDB.Begin()
		fmt.Println("解析到英雄数据:",len(heros))
		for i := 0; i <= len(heros)-1; i++ {
			_, err := db_mysql.SaveHero(heros[i])
			if err != nil {
				fmt.Println(err.Error())
				affair.Rollback()
			}
		}
		affair.Commit()
	}
	fmt.Printf("数据库当中有英雄数据%d条\n", HeroCount)

	http.HandleFunc("/",reqhandlet.Index)

	http.HandleFunc("/log",reqhandlet.Login)

	http.HandleFunc("/home.html",reqhandlet.Home)

	http.HandleFunc("/table.html",reqhandlet.Table)


	err = http.ListenAndServe(":9092", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
