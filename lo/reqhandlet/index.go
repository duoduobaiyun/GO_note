package reqhandlet


import (
	"html/template"
	"net/http"
)

func Index(writer http.ResponseWriter,request *http.Request) {
	tempt,err:=template.ParseFiles("./vie/index.html")
	if err != nil {
		errorTempt,_ := template.ParseFiles("./vie/err.html")
		errorTempt.Execute(writer,err.Error())
		return
	}
	tempt.Execute(writer,nil)
}