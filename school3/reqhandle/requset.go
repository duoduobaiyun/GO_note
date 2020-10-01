package reqhandle

import (
	"html/template"
	"net/http"
)

func Index(writer http.ResponseWriter,request *http.Request){
	//writer.Write([]byte("欢迎进入豆瓣信息采集系统"))
	temp,err := template.ParseFiles("./view/index.html")
	if err != nil {
		errorTemp,_ :=template.ParseFiles("./view/error.html")
		errorTemp.Execute(writer,err.Error())
		return
	}

	temp.Execute(writer,nil)


}
