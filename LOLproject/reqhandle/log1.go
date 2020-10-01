package reqhandle

import (
	"LOLproject/mysql_db"
	"LOLproject/object"
	"fmt"
	"html/template"
	"net/http"
)

func Log1(writer http.ResponseWriter, request *http.Request)  {
	err := request.ParseForm()
	if err != nil {
		errTempt, _ := template.ParseFiles("./views/error.html")
		errTempt.Execute(writer, "账号登录，解析参数错误,请重试")
		return
	}
	userName := request.FormValue("user_name")
	userPwd := request.FormValue("user_pwd")
	if userName == "" || userPwd == "" {
		errTmp, _ := template.ParseFiles("./views/error.html")
		errTmp.Execute(writer, "用户名或密码为空，请检查后重试")
		return
	}
	admin, err := mysql_db.Admin(userName, userPwd)
	if err != nil {
		fmt.Println(err.Error())
		errTmp, _ := template.ParseFiles("./views/error.html")
		errTmp.Execute(writer, "用户名或密码错误，请检查后重试")
		return
	}
	resTmp, err := template.ParseFiles("./views/response.html")
	if err != nil {
		errTmp, _ := template.ParseFiles("./views/error.html")
		errTmp.Execute(writer, "查询数据发生错误，请稍后再试")
		return
	}
	heros, err := mysql_db.SelectAll()
	if err != nil {
		errTmp, _ := template.ParseFiles("./views/error.html")
		errTmp.Execute(writer, "查询数据发生错误，请稍后重试")
		return
	}
	res := object.Response{
		admin.Username,
		heros,
	}
	resTmp.Execute(writer, res)
}
