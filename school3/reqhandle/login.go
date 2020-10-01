package reqhandle

import (
	"html/template"
	"net/http"
	"school3/entity"
	"school3/spider_db"
)

func Login(writer http.ResponseWriter,request *http.Request)  {
	err :=request.ParseForm()
	if err!=nil {
		errTemp , _:=template.ParseFiles("./view/error.html")

		errTemp.Execute(writer,"账号错误")
		return
	}
	username := request.FormValue("username")
	userpwd := request.FormValue("password")
	if username==""||userpwd=="" {
		errTmp,_ := template.ParseFiles("./view/error.html")
		errTmp.Execute(writer, "用户名或密码为空请输入")
		return
	}

	admin,err:= spider_db.QueryAdmin(username, userpwd)
	//fmt.Println(err)
	if err != nil {
		errTmp,_ := template.ParseFiles("./view/error.html")
		errTmp.Execute(writer, "用户名或密码错误")

	}


	homeTemp ,err := template.ParseFiles("./view/home.html")
	if err != nil {
		errTmp,_ := template.ParseFiles("./view/error.html")
		errTmp.Execute(writer, "程序查询异常，请稍后再试")
		return

	}
	//
	movies,err:=spider_db.QueryAllMovies()
	if err!=nil {
		errTmp,_ := template.ParseFiles("./view/error.html")
		errTmp.Execute(writer,"查询数据发生错误，请稍后再试")
		return
	}
	//

	homeDate:=entity.HomeDate{
		admin.Name,
		movies,
	}

	homeTemp.Execute(writer,homeDate)

	//fmt.Println(admin==nil)
	//fmt.Println(homeTemp==nil)

	//homeTemp.Execute(writer,admin.Name)

}
