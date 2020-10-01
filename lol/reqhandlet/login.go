package reqhandlet

import (
	"fmt"
	"html/template"
	"net/http"
	"project1/LoL/object"
	"project1/LoL/spider_lol"
)

func Login(writer http.ResponseWriter,request *http.Request)  {
	err :=request.ParseForm()
	if err != nil {
		if err != nil {
			errTempt,_ := template.ParseFiles("./vie/error.html")
			errTempt.Execute(writer,"账号登录，解析参数错误,请重试")
			return
		}
	}

	userName := request.FormValue("user_name")
	userPwd := request.FormValue("user_pwd")

	if userName == "" || userPwd == ""{
		errTmp,_ := template.ParseFiles("./vie/error.html")
		errTmp.Execute(writer,"用户名密码没有填写，请重试")
		return
	}

	query, err :=spider_lol.QueryAdmin(userName,userPwd)
	if err != nil {
		fmt.Println(err.Error())
		errTmp,_ := template.ParseFiles("./vie/error.html")
		errTmp.Execute(writer,"用户名或密码错误，请检查后重试")
		return
	}

	homePipe,err:= template.ParseFiles("./vie/home.html","")
	if err != nil {
		errTmp,_ := template.ParseFiles("./vie/error.html")
		errTmp.Execute(writer,"查询数据发生错误，请稍后再试")
		return
	}

	heros,err:=spider_lol.QueryAllHero()
	if err != nil {
		errTmp,_ := template.ParseFiles("./vie/error.html")
		errTmp.Execute(writer,"查询数据发生错误，请稍后再试")
		return
	}

	homeData :=object.HomeData{
		AdminName: query.Name,
		Heros:    heros,
	}


	homePipe.Execute(writer,homeData)

}

