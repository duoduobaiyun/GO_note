package handfunc

import (
	"html/template"
	"net/http"
)

func Index(writer http.ResponseWriter,request *http.Request )  {
	tempt,err:=template.ParseFiles("./view/index.html")
	if err!=nil {
		errortempt,_:=template.ParseFiles("./view/error/error.html")
		errortempt.Execute(writer,nil)
		return
	}
	tempt.Execute(writer,nil)
}