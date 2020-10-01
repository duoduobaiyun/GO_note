package handfunc

import (

	"fmt"
	"html/template"
	"lol(new)/db_mysql"
	"lol(new)/object1"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	//1.请求解析的参数
	err := request.ParseForm()
	if err != nil {
		errtempt, _ := template.ParseFiles("./view/error/error.html")
		errtempt.Execute(writer, nil)
		return
	}
	userName := request.FormValue("name")
	userPwd := request.FormValue("password")
	//查询用户是否存在 fmt.Println(userName,userPwd)
	admin, err := db_mysql.QueryAdmin(userName, userPwd)
	fmt.Println(admin)
	if err != nil {
		errtempt, _ := template.ParseFiles("./view/error/error.html")
		errtempt.Execute(writer, nil)
		return
	}
	//根据查询的结果，加载模板文件
	hometempt,err:=template.ParseFiles("./view/homepage.html")
	if err != nil {
		errtempt, _ := template.ParseFiles("./view/error/error.html")
		errtempt.Execute(writer, nil)
		return
	}
	//4.查询数据库表，获取所有的需展示的数信息
	heroes,err:=db_mysql.QueryAllHero()
	if err != nil {
		errtempt, _ := template.ParseFiles("./view/error/error.html")
		errtempt.Execute(writer, nil)
		return
	}
	fmt.Println(heroes)
	heroData:=object1.HeroDate{
		AdminName:admin.Name,
		Hero:     heroes,
	}
	hometempt.Execute(writer,heroData)
}