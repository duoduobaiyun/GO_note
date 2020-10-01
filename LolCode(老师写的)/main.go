package main

import (
	"fmt"
	"net/http"
	"os"
	"LolCode/controller"
	"LolCode/util"
)

func main() {
	fmt.Println("欢迎进入英雄联盟英雄信息系统")

	//打开数据库
	err := util.OpenDb()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
		return
	}

	//关闭数据库
	defer util.CloseDb()

	//首页
	http.HandleFunc("/", controller.Index)

	//采集英雄数据
	http.HandleFunc("/collect", controller.Collect)

	//上一页
	http.HandleFunc("/pre", controller.PrePage)
	//下一页
	http.HandleFunc("/next", controller.NextPage)

	//删除
	http.HandleFunc("/delete",controller.Del)

	err = http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
		return
	}
	fmt.Println("退出英雄联盟英雄信息系统，欢迎下次光临")
}
