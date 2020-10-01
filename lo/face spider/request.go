package face_spider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lo/db_mysql"
	"net/http"
	"os"
)

func RequestFace() {
	err := db_mysql.OPenDB()
	if err != nil {
		fmt.Println("打开数据库遇到错误：", err.Error())
		os.Exit(0)
	}

	defer db_mysql.CloseDB()

	HeroCount, err := db_mysql.QueryHeroCount()
	if err != nil {
		fmt.Println(err.Error())
	}
	if HeroCount <= 0 {
		url := "https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js"

		client := http.Client{}

		request, err := http.NewRequest("GET", url, nil) //获取
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
		var herolist object.HeroList
		err = json.Unmarshal(heroData, &herolist)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}