package controller

import (
	"net/http"
	"LolCode/util"
	"LolCode/bean"
	"LolCode/dao"
	"html/template"
	"io/ioutil"
	"encoding/json"
)

func Collect(writer http.ResponseWriter, request *http.Request) {
	count, err := dao.QueryHeroCount()

	if err != nil {
		util.ErrorMsg(writer, err)
		return
	}
	if count <= 0 {
		response, err := http.Get("https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js")
		if err != nil {
			util.ErrorMsg(writer, err)
			return
		}
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			util.ErrorMsg(writer, err)
			return
		}
		defer response.Body.Close()

		var heroList bean.HeroList
		err = json.Unmarshal(bytes, &heroList)
		if err != nil {
			util.ErrorMsg(writer, err)
			return
		}

		//保存数据
		tx, err := util.LolDB.Begin()
		if err != nil {
			util.ErrorMsg(writer, err)
			return
		}

		for _, value := range heroList.Hero {
			_, err := dao.SaveHero(value)
			if err != nil {
				util.ErrorMsg(writer, err)
				tx.Rollback()
				return
			}
		}
		tx.Commit()
	}

	home, err := template.ParseFiles("./static/index.html")
	if err != nil {
		util.ErrorMsg(writer, err)
		return
	}

	indexData := bean.IndexData{}
	indexData.IsNeedSpiderData = count <= 0

	if count > 0 {
		indexData.PageNum = count/PER_NUMER + 1
		indexData.CurrentNum = 1
		indexData.IsPageButton = indexData.PageNum > 1
		heros, err := dao.QueryHeros(0, PER_NUMER)
		if err != nil {
			util.ErrorMsg(writer, err)
			return
		}
		indexData.Heros = heros
	}
	home.Execute(writer, indexData)
}
