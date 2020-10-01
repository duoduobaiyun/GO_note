package util

import (
	"html/template"
	"net/http"
)

func ErrorMsg(writer http.ResponseWriter, err error) {
	errHtml, _ := template.ParseFiles("./static/error.html")
	errHtml.Execute(writer, err.Error())
}
