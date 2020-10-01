package reqhandlet

import (
	"fmt"
	"html/template"
	"net/http"
	"project1/lo/db_mysql"
	"project1/lo/object"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("login....")
	err := request.ParseForm()
	if err != nil {
		errTempt, _ := template.ParseFiles("./vie/err.html")
		errTempt.Execute(writer, "账号登录，解析参数错误,请重试")
		return
	}

	userName := request.FormValue("user_name")
	userPwd := request.FormValue("user_pwd")

	if userName == "" || userPwd == "" {
		errTmp, _ := template.ParseFiles("./vie/err.html")
		errTmp.Execute(writer, "用户名密码没有填写，请重试")
		return
	}

	query, err := db_mysql.QueryAdmin(userName, userPwd)
	if err != nil {
		fmt.Println(err.Error())
		errTmp, _ := template.ParseFiles("./vie/err.html")
		errTmp.Execute(writer, "用户名或密码错误，请检查后重试")
		return
	}

	mainPage, err := template.ParseFiles("./vie/main.html")
	if err != nil {
		errTmp, _ := template.ParseFiles("./vie/err.html")
		errTmp.Execute(writer, "查询数据发生错误，请稍后再试1")
		return
	}

	heroes, err := db_mysql.QueryAllHero()
	if err != nil {
		errTmp, _ := template.ParseFiles("./vie/err.html")
		errTmp.Execute(writer, "查询数据发生错误，请稍后再试")
		return
	}

	fmt.Println(len(heroes))
	homeData := object.HomeData{
		AdminName: query,
		Heroes:    heroes,
	}

	mainPage.Execute(writer, homeData)
}
