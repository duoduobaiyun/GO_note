package reqhandle

import (
	"fmt"
	"html/template"
	"net/http"
)



func Index(writer http.ResponseWriter,request *http.Request){
	//writer.Write([]byte("欢迎进入豆瓣信息采集系统"))

	fmt.Println(template.ParseFiles("./error.html"))

	temp,err := template.ParseFiles("./index.html")
	if err != nil {
		errorTemp,_ :=template.ParseFiles("./error.html")
		errorTemp.Execute(writer,err.Error())
		return
	}


	temp.Execute(writer,nil)


}
