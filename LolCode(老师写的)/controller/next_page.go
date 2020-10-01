package controller

import (
	"net/http"
	"strconv"
	"LolCode/util"
	"LolCode/dao"
	"LolCode/bean"
	"html/template"
)

/**
 * 下一页
 */
func NextPage(writer http.ResponseWriter, request *http.Request) {
	pageStr := request.FormValue("page_num")
	pageNum, err := strconv.Atoi(pageStr)
	if err != nil {
		util.ErrorMsg(writer, err)
		return
	}

	count, err := dao.QueryHeroCount()
	if err != nil {
		util.ErrorMsg(writer, err)
		return
	}

	pageNum++
	if pageNum >= int(count)/PER_NUMER+1 {
		pageNum = int(count)/PER_NUMER + 1
	}

	start := pageNum*PER_NUMER
	if start >= int(count) {
		start = int(count)
	}
	limit := 0
	if int(count)-start > PER_NUMER {
		limit = PER_NUMER
	} else {
		limit = int(count) - start
	}
	heros, err := dao.QueryHeros(start, limit)
	if err != nil {
		util.ErrorMsg(writer, err)
		return
	}

	indexData := bean.IndexData{}
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
