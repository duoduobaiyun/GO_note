package handle

import (
	"html/template"
	"net/http"
)

func Index(writer http.ResponseWriter , request *http.Request)  {
    tempt,err:=template.ParseFiles("./view/index.html")
	if err != nil {
		errorTempt,_:=template.ParseFiles("./view/error.html")
		errorTempt.Execute(writer,err.Error())
		return
	}
	tempt.Execute(writer,nil)
}
