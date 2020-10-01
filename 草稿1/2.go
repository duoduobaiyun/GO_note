package main

//加载mysql驱动
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	/**
	  需求：在go语言文件中，通过使用go编程语法，完成对数据库的连接（mysql数据库）
	  数据库软件：Oracle、SqlServer、MySQL ...多款产品
	  数据库驱动：代码库，公开的
	  步骤：
		* 1、准备数据库驱动程序：MySQL
	    * 草稿2、执行数据库驱动程序（查看MySQL驱动说明文档）
		* 3、连接数据库：成功、失败
		* 4、操作数据库（数据库表创建、数据的添加、删除。。。。)
		* 5、关闭数据库连接close


	  电脑：
	     1、鼠标（蓝牙鼠标）：安装驱动
		 草稿2、键盘：插入到usb插口，正在安装驱动

	   转换器、转换头

	*/


	//3、连接数据库
	/**
	  driverName：驱动名称，mysql
	*/
	//"root:yu123456@tcp(127.0.0.1:3306)/ruanda?charset=utf8"
	connStr := "root:yu123456@tcp(127.0.0.1:3306)/ruanda?charset=utf8"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		fmt.Println("连接失败")
		log.Fatal(err)
		return
	}
	fmt.Println("连接成功")
	db.Close()//关闭数据库

}

