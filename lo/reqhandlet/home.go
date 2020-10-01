package reqhandlet

import (
	"fmt"
	"html/template"
	"net/http"
	"project1/lo/db_mysql"
	"project1/lo/object"
)

func Home(writer http.ResponseWriter, request *http.Request)  {
	homePage, err := template.ParseFiles("./vie/home.html")
	if err!=nil {
		errTmp, _ := template.ParseFiles("./vie/err.html")
		errTmp.Execute(writer, "查询数据发生错误，请稍后再试1")
		return
	}


	heroes, err := db_mysql.QueryAllHero()
	if err != nil {
		errTmp, _ := template.ParseFiles("./vie/err.html")
		errTmp.Execute(writer, "查询数据发生错误，请稍后再试")
		return
	}


	fmt.Println(len(heroes))
	homeData := object.HomeData{
		Heroes:heroes,
	}

	homePage.Execute(writer,homeData )
}
