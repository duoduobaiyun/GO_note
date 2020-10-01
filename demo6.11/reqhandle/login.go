package reqhandle

import (
	"DoubanNewMovie/spider_db"
	"fmt"
	"html/template"
	"net/http"
)

/**
 * 处理管理员的登录请求
 */
func Login(writer http.ResponseWriter,request *http.Request){
	//1、解析请求参数
	err :=request.ParseForm()//解析参数
	if err != nil {
		errTempt,_ := template.ParseFiles("./error.html")
		//200、404 系统
		// ERROR_NUM：自定义系统的错误代码
		// 10001 ： 用户名错误
		// 10002 : 用户名不存在
		// 10003 : 账号登录，解析参数错误
		errTempt.Execute(writer,"账号登录，解析参数错误,请重试")
		return
	}

	userName := request.FormValue("user_name")//管理员填写的登录账号
	userPwd := request.FormValue("user_pwd")//管理员填写的登录密码
	//2、检查参数
	if userName == "" || userPwd == ""{
		errTmp,_ := template.ParseFiles("./error.html")
		errTmp.Execute(writer,"用户名或密码为空，请检查后重试")
		return
	}

	//3、查询数据库
	admin, err :=spider_db.QueryAdmin(userName,userPwd)
	if err != nil {
		fmt.Println(err.Error())
		errTmp,_ := template.ParseFiles("./error.html")
		errTmp.Execute(writer,"用户名或密码错误，请检查后重试")
		return
	}

	//4、根据查询的结果，加载模板文件
	//管理主界面： 欢迎你：xxxx
	homeTmp,err := template.ParseFiles("./home.html")
	if err != nil {
		errTmp,_ := template.ParseFiles("./error.html")
		errTmp.Execute(writer,"查询数据发生错误，请稍后再试")
		return
	}
	homeTmp.Execute(writer,admin.Name)
}
