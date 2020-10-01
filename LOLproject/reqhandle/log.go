package reqhandle

import (
	"html/template"
	"net/http"
)

func Log(writer http.ResponseWriter, request *http.Request)  {
	temp, err := template.ParseFiles("./views/login.html")
	if err != nil {
		errtemp, _ := template.ParseFiles("./views/error.html")
		errtemp.Execute(writer, err.Error())
		return
	}
	temp.Execute(writer, nil)
}
