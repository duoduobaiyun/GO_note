package handfunc

import (
	"LeagueOfLegends/mysql_db"
	"LeagueOfLegends/object"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func UserRegistration (writer http.ResponseWriter, request *http.Request) {
	tempt, err := template.ParseFiles("./views/userregistration/userregistration.html")
	if err != nil {
		errtempt, _ := template.ParseFiles("./views/error/error.html")
		errtempt.Execute(writer,nil)
		return
	}
	tempt.Execute(writer,nil)

}


//解析注册页面的数据，并存入数据库中
func ParseSI(writer http.ResponseWriter, request *http.Request) {
	//1.解析请求的参数
	err := request.ParseForm()
	if err != nil {
		errTempt, _ := template.ParseFiles("./views/error/error.html")
		errTempt.Execute(writer,nil)
		return
	}

	//2.获取网页的基本信息
	name := request.FormValue("user_name")
	email := request.FormValue("user_email")
	pwd := request.FormValue("user_password")
	Apwd := request.FormValue("user_confirm_password")
	sex := request.FormValue("user_sex")
	phone := request.FormValue("user_phone1")
	//fmt.Println(name,email,pwd,Apwd,sex,phone)

	//3.比对 pwd是否与Apwd相同，不相同则结束，并返回密码不相同
	if pwd != Apwd {
		result := "两次密码不相同"
		writer.Write([]byte(result))
		return
	}

	//4.生成新用户的username（共12...位）2019+随机数（6位）+id（01,100.。。。。。）
	//4.1获取admin_id用于设计username
	count, err := mysql_db.QueryAdminCount()
	if err != nil {
		fmt.Println(err.Error())
		errTempt, _ := template.ParseFiles("./views/error/error.html")
		errTempt.Execute(writer,nil)
		return
	}
	//4.2生成随机数
	str := fmt.Sprintf("%03v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	username := "2019"+ str + "0" + fmt.Sprintf("%v",count + 1)
	//fmt.Println(count,username)

	//5，将数据存入数据库中
	var admin object.Admin
	admin = object.Admin{
		Id : count + 1,
		Name : name,
		Email : email,
		UserName : username,
		UserPwd : pwd,
		Phone : phone,
		Authority : 1,
		Sex : sex,
	}
	v,err := template.ParseFiles("./views/userregistration/parsesi.html")
	if err != nil {
		errTempt, _ := template.ParseFiles("./views/error/error.html")
		errTempt.Execute(writer,nil)
		return
	}
	v.Execute(writer,admin)

	mysql_db.SavieAdminInfo2Db(admin)
}
