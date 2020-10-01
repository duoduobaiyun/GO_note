package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


func main() {
	connser:="root:123@tcp(127.0.0.1:3306)/wish?charset=utf8"
	db,err:=sql.Open("mysql",connser)
	if err!=nil {
		log.Fatal(err)
		fmt.Println("连接失败")
	}

	result,err:=db.Query("select * from student")
	if err!=nil {
		log.Fatal(err)
		fmt.Println("数据库查询失败")
	}
	//打印出字段名,但是也可以不写
	columns,err:=result.Columns()
	if err!=nil {
		log.Fatal(err)
		fmt.Println("查询失败")
	}
	fmt.Println(columns)

	for result.Next(){
		var id  int
		var name string
		var sex  string
		var grade int
		err:=result.Scan(&id,&name,&sex,&grade)
		if err!=nil {
			log.Fatal(err)
			break
		}
		fmt.Println(id,name,sex,grade)
	}
}
