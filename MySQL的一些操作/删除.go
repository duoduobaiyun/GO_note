package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	connstr:="root:123@tcp(127.0.0.1:3306)/wish?charset=utf8"
	db,err:=sql.Open("mysql",connstr)
	if err!=nil {
		log.Fatal(err)
		fmt.Println("连接成功")
	}

	result,err:=db.Exec("delete from student where id=561")
	if err!=nil {
		log.Fatal(err)
		fmt.Println("删除失败")
	}
	fmt.Println(result.RowsAffected())
}
