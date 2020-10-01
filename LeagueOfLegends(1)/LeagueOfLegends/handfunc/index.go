package handfunc

import (
	"html/template"
	"net/http"
)

func Index (writer http.ResponseWriter, request *http.Request) {
	tempt, err := template.ParseFiles("./views/index.html")
	if err != nil {
		errtempt, _ := template.ParseFiles("./views/error/error.html")
		errtempt.Execute(writer,nil)
		return
	}
	tempt.Execute(writer,nil)
}
