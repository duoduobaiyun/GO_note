package controller

import (
	"net/http"
	"html/template"
	"LolCode/dao"
	"LolCode/util"
	"LolCode/bean"
)

const PER_NUMER = 10;

func Index(writer http.ResponseWriter, req *http.Request) {
	index, err := template.ParseFiles("./static/index.html")
	if err != nil {
		util.ErrorMsg(writer, err)
		return
	}

	count, err := dao.QueryHeroCount()

	if err != nil {
		util.ErrorMsg(writer, err)
		return
	}

	indexData := bean.IndexData{}

	indexData.IsNeedSpiderData = count <= 0

	if count > 0 {
		indexData.PageNum = count/PER_NUMER + 1
		indexData.IsPageButton = indexData.PageNum > 1
		indexData.CurrentNum = 1
		allHeros, err := dao.QueryHeros(0, PER_NUMER)
		if err != nil {
			util.ErrorMsg(writer, err)
			return
		}
		indexData.Heros = allHeros
	}
	index.Execute(writer, indexData)
}
