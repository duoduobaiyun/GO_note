package handfunc

import (
	"LeagueOfLegends/mysql_db"
	"LeagueOfLegends/object"
	"fmt"
	"github.com/skip2/go-qrcode"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var Str string
var adminData object.Admin

//页面1
func ForgotPasswprd(writer http.ResponseWriter, request *http.Request) {
	//加载forgotpassword模板
	tempt, err := template.ParseFiles("./views/forgotPassword/forgotPassword.html")
	if err != nil {
		errtempt, _ := template.ParseFiles("./views/error/error.html")
		errtempt.Execute(writer, nil)
		return
	}
	tempt.Execute(writer, nil)

}

//页面3
func YZM(writer http.ResponseWriter, request *http.Request) {
	//1.解析请求的参数
	err := request.ParseForm()
	if err != nil {
		errTempt, _ := template.ParseFiles("./views/error/error.html")
		errTempt.Execute(writer, nil)
		return
	}

	//2.获取网页的基本信息
	code := request.FormValue("code")
	//fmt.Println(code,   Str)
	//页面code与生成的验证码对比
	if code != Str {
		errTempt, _ := template.ParseFiles("./views/error/error.html")
		errTempt.Execute(writer, nil)
		return
	}

	//4.修改MySQL数据库中的密码,加载相应的页面
	mysql_db.UpdateAdminInfo2Db(adminData.UserPwd, adminData.UserName)


	v, err := template.ParseFiles("./views/forgotPassword/captcha.html")
	if err != nil {
		errTempt, _ := template.ParseFiles("./views/error/error.html")
		errTempt.Execute(writer, nil)
		return
	}
	v.Execute(writer, adminData)

}

//页面2，获取验证码并刷新页面
func GetQrCode(writer http.ResponseWriter, request *http.Request) {

	Str = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)) //随机生成二维码
	err := qrcode.WriteFile(Str, qrcode.Medium, 256, "./static/qrcode/YZM1.png")
	if err != nil {
		tempErr, err := template.ParseFiles("./views/error/error.html")
		tempErr.Execute(writer, err.Error())
		fmt.Println("生成验证码失败")
		return
	}
	//验证码生成成功, 刷新页面
	forgetPwdPage, err := template.ParseFiles("./views/forgotPassword/GetNewcode.html")
	if err != nil {
		tempErr, err := template.ParseFiles("./views/error/error.html")
		tempErr.Execute(writer, err.Error())
		return
	}
	//fmt.Println("新生成的验证码:", Str)
	forgetPwdPage.Execute(writer, nil)
}

//页面2，加载页面模板，获取解析页面1的值，存入admin对象中
func GetNewcode(writer http.ResponseWriter, request *http.Request) {
	//生成一张验证码
	Str = "先点击获取验证码，再扫码获取具体的验证码"
	err := qrcode.WriteFile(Str, qrcode.Medium, 256, "./static/qrcode/YZM1.png")
	if err != nil {
		tempErr, err := template.ParseFiles("./views/error/error.html")
		tempErr.Execute(writer, err.Error())
		fmt.Println("生成验证码失败")
		return
	}

	//解析请求的参数
	err = request.ParseForm()
	if err != nil {
		errTempt, _ := template.ParseFiles("./views/error/error.html")
		errTempt.Execute(writer, nil)
		return
	}

	//获取网页的基本信息
	userName := request.FormValue("userName")
	phone := request.FormValue("phone")
	newPwd := request.FormValue("newPwd")
	email := request.FormValue("user_email")

	//查询admin的基本信息并比对，后加载相应的模板文件
	admin, err := mysql_db.QueryYZMAdmin(userName, phone, email)
	if err != nil {
		fmt.Println(err.Error())
		errTempt, _ := template.ParseFiles("./views/error/error.html")
		errTempt.Execute(writer, nil)
		return
	}

	adminData = object.Admin{
		Name:     admin.Name,
		UserName: userName,
		Phone:    phone,
		Email:    email,
		UserPwd : newPwd,
	}
	v, err := template.ParseFiles("./views/forgotPassword/getNewcode.html")
	if err != nil {
		errTempt, _ := template.ParseFiles("./views/error/error.html")
		errTempt.Execute(writer, nil)
		return
	}
	v.Execute(writer, nil)
}

func GetPhoneCode(){}


