package controller

import (
	"net/http"
	"LolCode/util"
	"LolCode/dao"
	"LolCode/bean"
	"strconv"
	"html/template"
)

/**
 * 上一页
 */
func PrePage(writer http.ResponseWriter, request *http.Request) {
	pageStr := request.FormValue("page_num")
	pageNum, err := strconv.Atoi(pageStr)
	if err != nil {
		util.ErrorMsg(writer, err)
		return
	}

	start := pageNum*PER_NUMER - 10
	if start <= 0 {
		start = 0
	}

	heros, err := dao.QueryHeros(start, PER_NUMER)
	if err != nil {
		util.ErrorMsg(writer, err)
		return
	}

	indexData := bean.IndexData{}
	pageNum = pageNum - 1
	if pageNum <= 1 {
		pageNum = 1
	}
	indexData.CurrentNum = pageNum
	indexData.IsNeedSpiderData = false //有上一页下一页，就不用采集数据，因此为false
	indexData.IsPageButton = true
	indexData.Heros = heros

	home, err := template.ParseFiles("./static/index.html")
	if err != nil {
		util.ErrorMsg(writer, err)
		return
	}

	home.Execute(writer, indexData)
}
