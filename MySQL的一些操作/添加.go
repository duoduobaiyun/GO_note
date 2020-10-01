package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	connstr:="root:123@tcp(127.0.0.1)/wish?charset=utf8"
	db,err:=sql.Open("mysql",connstr)
	if err!=nil {
		log.Fatal(err)
		fmt.Println("连接失败")
	}
	result,err:=db.Exec("insert into student (id,name ,sex,grade) values (600,'王鹏','女',110) ")
	if err!=nil {
		log.Fatal(err)
		fmt.Println("添加失败")
	}
	fmt.Println(result)
   // fmt.Println(result.RowsAffected())

}
