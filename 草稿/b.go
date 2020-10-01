package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"unicode/utf8"
)

func main() {
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		username:=request.FormValue("username")
		length:=utf8.RuneCountInString(username)//字符串变成int类型，张三  草稿2
		if length > 8 {
			writer.Write([]byte("名字不能超过8位"))
			fmt.Println("名字不能超过8位")
			return
		}
		fmt.Println("用户名:",username)

		password:=request.FormValue("password")
		PasResult,err:=regexp.MatchString(`^[0-9]+$`,password)
		if err!=nil {
			log.Fatal()
		}
		if !PasResult {
			fmt.Println("格式不正确")
		}
		fmt.Println("密码:",password)

		XueLi:=request.FormValue("XueLi")
		XueLiSlice:=[]string{"xiaoxue","chuzhong","gaozhong","dazhuang"}
		//var xueliResult bool
		for _,value:= range XueLiSlice {
			if value != XueLi {
			//	//xueliResult = true
            //   break
			//}else {
			//	//xueliResult =false
			//	return
			}
		}
		//   if !xueliResult  {
		//	writer.Write([]byte("学历不合法，请重新输入"))
		//	fmt.Println("学习不合法")
		//}
		fmt.Println("学历:",XueLi)

		writer.Write([]byte("恭喜，你成功了"))
	})
	err:=http.ListenAndServe("127.0.0.1:9000",nil)
	if err!=nil {
		log.Fatal(err)
	}
}
