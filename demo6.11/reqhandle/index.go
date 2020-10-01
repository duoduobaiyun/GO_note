package reqhandle

import (
	"html/template"
	"net/http"
)

/**
 * 处理默认的首页
 */
func Index(writer http.ResponseWriter,request *http.Request){

	//html：div + css + js ....
	//index.html文件：div+css 模板文件
	tempt,err := template.ParseFiles("./index.html")
	if err != nil {
		//给浏览器返回错误信息
		errorTempt,_ := template.ParseFiles("./error.html")
		errorTempt.Execute(writer,err.Error())
		return
	}
	//.exe : execute:执行
	tempt.Execute(writer,nil)//nil
}
