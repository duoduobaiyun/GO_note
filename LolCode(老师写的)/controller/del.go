package controller

import (
	"net/http"
	"strconv"
	"LolCode/util"
	"LolCode/dao"
	"errors"
	"fmt"
)

func Del(writer http.ResponseWriter, request *http.Request) {
	heroId := request.FormValue("hero_id")

	id, err := strconv.Atoi(heroId)

	if err != nil {
		util.ErrorMsg(writer, err)
		return
	}

	row, err := dao.DeleteHero(id)
	if err != nil {
		util.ErrorMsg(writer, err)
		return
	}

	if row <= 0 {
		util.ErrorMsg(writer, errors.New("未删除成功"))
		return
	}

	//刷新数据
	pageStr := request.FormValue("page_num")
	pageNum, err := strconv.Atoi(pageStr)
	if err != nil {
		util.ErrorMsg(writer, err)
		return
	}

	fmt.Println(pageNum)
}
