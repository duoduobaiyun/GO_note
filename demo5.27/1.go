package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	connStr:="root:123@tcp(127.0.0.1:3306)/college?charset=utf8"
	db,err:=sql.Open("mysql",connStr)
	if err!=nil {
		fmt.Println("连接失败")
		log.Fatal(err)
		return
	}
	fmt.Println("连接成功")
	db.Close()
}