package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"project1/LoL/object"
	"project1/LoL/parse"
	"project1/LoL/reqhandlet"
	"project1/LoL/spider_lol"
)

func main() {
	err:=spider_lol.OpenDB()
	if err!=nil{
		fmt.Println("打开数据库遇到错误：", err.Error())
		os.Exit(0)
	}
	
	HeroCount,err:=spider_lol.QueryHeroCount()
	if err != nil {
		fmt.Println(err.Error())
	}
	if HeroCount<=0 {
		url:="https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js"

		client:=http.Client{}

		request,err:=http.NewRequest("GET",url,nil)
		if err != nil {
			fmt.Println(request)
			return
		}
		resp,err:=client.Do(request)
		if err != nil {
			fmt.Println(request)
			return
		}
		heroData,err:=ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(request)
			return
		}
		fmt.Println(string(heroData))

		var herolist object.HeroList
		err =json.Unmarshal(heroData,&herolist)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		heros:=parse.Pagehtml(herolist.Hero)
		fmt.Println(heros)
		affair,err:=spider_lol.SpiderDb.Begin()
		for i := 0; i <= len(heros)-1; i++ {
			_,err:=spider_lol.SaveHero(heros[i])
			if err !=nil {
				fmt.Println(err.Error())
				affair.Rollback()
			}
		}
		affair.Commit()

		spider_lol.CloseDB()
		if err !=nil {
			fmt.Println(err.Error())
		}

	}
	fmt.Printf("数据库当中有电影数据%d条\n", HeroCount)

	http.HandleFunc("/", reqhandlet.Index)

	http.HandleFunc("/login", reqhandlet.Login)


	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err.Error())
	}





}
