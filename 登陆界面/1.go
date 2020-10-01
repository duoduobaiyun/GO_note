package main

import (
	"fmt"
	"net/http"
)

func main() {
	 /*
	     作业一：  ①  编写user_info.html，该页面用于实现用户的实名信息认证。
	     该页面上包含的信息有：真实姓名，性别（单选框），出生年月，家庭住址，身份证号，学历（下拉选择框），
	     发证机关（xxx公安局），该页面有两提交按钮，点击提交按钮，可以提交数据到服务器。
	     ② 使用Go语言程序编写一个服务器程序，在本地机上运行，接收user_info页面提交的数据。在服务器端解析前端提交的数据，并对数据做检查，最后根据检查返回提示信息。
	     数据规范：  姓名：最长不能超过12个字，名字不能为空，否则提示姓名不符合规范。  性别：有男、女两个选项，两个选项必须选择一个，否则提示：请选择性别信息。
	     出生年月：出生年月日不能超过2020年5月9日，否则提示：太着*/
	http.HandleFunc("/user_info",register)
	fmt.Println("正在监听")
	err:=http.ListenAndServe("127.0.0.1:9001",nil)
	if err!=nil {
		fmt.Println(err.Error())
	}
}

func register(writer http.ResponseWriter, request *http.Request)  {
      request.ParseForm()
      username:=request.FormValue("username")
	  if len(username) == 0  {
	  	fmt.Println("名字不能为空")
		  return
	  }else if len(username) <= 12 {
		 fmt.Println("名字最长不能超过12个字")
		  return
	  }
      password:=request.Form.Get("password")
      address:=request.Form.Get("address")
      birthday:=request.Form.Get("birthday")

	  IdCard:=request.Form.Get("IdCard")
	  IssuingAuthority:=request.Form.Get("IssuingAuthority")
	  education:=request.Form.Get("education")
	  sex:=request.Form.Get("sex")
	  if sex == "_" {
		fmt.Println("请选择性别")
		  return
	  }
	  fmt.Println(username,password,address,birthday,IdCard,IssuingAuthority,education,sex)
	writer.Write([]byte("欢迎回来英雄"))
	  fmt.Println("结束")
}