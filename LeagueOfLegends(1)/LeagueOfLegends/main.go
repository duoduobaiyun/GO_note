package main

import (
	"LeagueOfLegends/handfunc"
	"LeagueOfLegends/mysql_db"
	"LeagueOfLegends/network_request"
	"fmt"
	"net/http"
)

func main() {
	 A:= [3]bool{}
	 fmt.Println(A)


	fmt.Println("欢迎进入英雄联盟数据收集系统")

	url := "https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js"

	//1.打开数据库
	err := mysql_db.SqlOpen()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//2.关闭数据库，防止数据信息存储不进去
	defer mysql_db.SqlClose()

	//3.查询hero数量
	heroCount, err := mysql_db.QueryHeroCount()
	if err != nil {
		fmt.Println(err.Error())
	}
	if heroCount <= 0 {
		//4.获取并解析网络数据
		herolist, err := network_request.GetHerolist(url)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//5.打开事务，并存储每一条数据信息
		tx, err := mysql_db.HeroDb.Begin() //打开事务
		for i := 0; i < len(herolist.Hero); i++ {
			_, err := mysql_db.SavieHeroInfo2Db(herolist.Hero[i])
			if err != nil {
				fmt.Printf(err.Error())
				tx.Rollback() //事务回滚
			}
		}
		tx.Commit() //提交事务
	}

	fmt.Println("数据库中hero的数量为:", heroCount)
	//hero,_ := mysql_db.QueryAllHeros()
	//for i:=0;i<148;i++{
	//	fmt.Println(hero)
	//}
	//heroSlice,_  := mysql_db.QueryAllHeros()
	//fmt.Println(heroSlice)

	//6.启动浏览器服务器并在网页上展示
	http.HandleFunc("/", handfunc.Index)                            //登录界面
	http.HandleFunc("/login", handfunc.Login)                       //展示页面
	http.HandleFunc("/forgotPassword", handfunc.ForgotPasswprd)     //忘记密码页面（提交到getNewcode）
	http.HandleFunc("/getNewcode", handfunc.GetNewcode)				//忘记密码页面2（提交到yzm）
	http.HandleFunc("/yzm", handfunc.YZM)                           //忘记密码页面2
	http.HandleFunc("/userregistration", handfunc.UserRegistration) //新用户注册页面
	http.HandleFunc("/parsesi", handfunc.ParseSI)                   //与注册界面关联，既完成后加载注册界面
	http.HandleFunc("/get_qrcode", handfunc.GetQrCode)

	//静态文件服务器
	http.Handle("/static/qrcode/", http.StripPrefix("/static/qrcode/", http.FileServer(http.Dir("./static/qrcode/"))))
	http.Handle("/static/error/", http.StripPrefix("/static/error/", http.FileServer(http.Dir("./static/error/"))))


	fmt.Println("正在启动9090端口浏览器服务器")
	err = http.ListenAndServe(":9090", nil)
	
}

