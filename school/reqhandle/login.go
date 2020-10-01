package reqhandle

import (
	"html/template"
	"net/http"
	"school/spider_db"
)

func Login(writer http.ResponseWriter,request *http.Request)  {
	err :=request.ParseForm()
	if err!=nil {
		errTemp , _:=template.ParseFiles("./error.html")

		errTemp.Execute(writer,"账号错误")
		return
	}
	username := request.FormValue("username")
	userpwd := request.FormValue("password")
	if username==""||userpwd=="" {
		errTmp,_ := template.ParseFiles("./error.html")
		errTmp.Execute(writer, "用户名或密码为空请输入")
		return
	}

	admin,err:= spider_db.QueryAdmin(username, userpwd)
	if err != nil {
		errTmp,_ := template.ParseFiles("./error.html")
		errTmp.Execute(writer, "用户名或密码错误")

	}


	homeTemp ,err := template.ParseFiles("./home.html")
if err != nil {
errTmp,_ := template.ParseFiles("./error.html")
errTmp.Execute(writer, "程序查询异常，请稍后再试")
return

}
homeTemp.Execute(writer,admin.Name)
}
