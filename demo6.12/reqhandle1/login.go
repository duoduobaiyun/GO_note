package reqhandle1

import (
	"demo6.12/spider_db"
	"fmt"
	"html/template"
	"net/http"
)

func Login(writer http.ResponseWriter,request *http.Request)  {
	err:=request.ParseForm()//解析函数
	if err!=nil {
		errTempt,_:=template.ParseFiles("./error1.html")
		errTempt.Execute(writer,"账号登录，解析参数错误,请重试")
		return
	}

	userName:=request.FormValue("username")
	userPwd:=request.FormValue("password")
	if userName == "" || userPwd == ""{
		errTmp,_ := template.ParseFiles("./error1.html")
		errTmp.Execute(writer,"用户名或密码为空，请检查后重试")
		return
	}

	admin, err :=spider_db.QueryAdmin(userName,userPwd)
	if err != nil {
		fmt.Println(err.Error())
		errTmp,_ := template.ParseFiles("./error1.html")
		errTmp.Execute(writer,"用户名或密码错误，请检查后重试")
		return
	}
	//4、根据查询的结果，加载模板文件
	//管理主界面： 欢迎你：xxxx
	homeTmp,err := template.ParseFiles("./home.html")
	if err != nil {
		errTmp,_ := template.ParseFiles("./error1.html")
		errTmp.Execute(writer,"查询数据发生错误，请稍后再试")
		return
	}
	homeTmp.Execute(writer,admin.Name)
}
