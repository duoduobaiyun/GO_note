package handfunc

import (
	"html/template"
	"lol(new)/db_mysql"
	"lol(new)/object1"
	"github.com/go-qrcode-master/bitset"
	"net/http"
)

var Str string
var adminDate  object1.HeroDate

func ForgetPwd(writer http.ResponseWriter, request *http.Request)  {
	//加载forgetPwd文件
	tempt,err:=template.ParseFiles("./view/forgetpasswed/gorgetpassword.html")
	if err!=nil {
		errtempt,_:=template.ParseFiles("./view/error/error.html")
		errtempt.Execute(writer,nil)
	}
	tempt.Execute(writer,nil)
}

func YZ(writer http.ResponseWriter,request *http.Request)  {
	//解析请求参数
	err:=request.ParseForm()
	if err!=nil {
		errtempt,_:=template.ParseFiles("./view/error/error.html")
		errtempt.Execute(writer,nil)
		return
	}

	//获取网页的基本信息
	code:=request.FormValue("code")
	if code!=Str {
		errtempt,_:=template.ParseFiles("./view/error/error.html")
		errtempt.Execute(writer,nil)
		return
	}
	//4.修改MySQL数据库中的密码,加载相应的页面
	db_mysql.UpdateAdminInfo2DB(adminDate)
}