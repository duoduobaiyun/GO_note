package handfunc

import (
	"LeagueOfLegends/mysql_db"
	"LeagueOfLegends/object"
	"fmt"
	"html/template"
	"net/http"
)

func Login (writer http.ResponseWriter, request *http.Request) {
	//1.解析请求的参数
	err := request.ParseForm()
	if err != nil {
		errTempt, _ := template.ParseFiles("./views/error/error.html")
		errTempt.Execute(writer,nil)
		return
	}

	userName := request.FormValue("name")
	userPwd := request.FormValue("password")

	//fmt.Println(userName,userPwd)
	//2.查询数据库，比对该用户是否存在
	admin, err := mysql_db.QueryAdmin(userName,userPwd)
	fmt.Println(admin)
	if err != nil {
		errTempt, _ := template.ParseFiles("./views/error/error.html")
		errTempt.Execute(writer,nil)
		return
	}

	//3.根据查询的结果，加载模板文件
	//homeTempt, err := template.ParseFiles("./views/home/homePage.html")
	homeTempt, err := template.ParseFiles("./views/home/homePage2.html")
	if err != nil {
		errTempt, _ := template.ParseFiles("./views/error/error.html")
		errTempt.Execute(writer,nil)
		return
	}

	//4.查询数据库表，获取所有的需展示的数信息
	heros, err := mysql_db.QueryAllHeros()
	if err != nil {
		errTempt, _ := template.ParseFiles("./views/error/error.html")
		errTempt.Execute(writer,nil)
		return
	}
	fmt.Println(heros)
	heroData := object.HeroData{
		AdminName: admin.Name,
		Hero:     heros,

	}
	homeTempt.Execute(writer,heroData)
}
