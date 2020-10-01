package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	connstr:="root:123@tcp(127.0.0.1:3306)/school?charset=utf8"
	db,err:=sql.Open("mysql",connstr)
	defer db.Close()
	if err!=nil {
		log.Fatal(err)
		fmt.Println("连接失败")
	}
	fmt.Println("连接成功")


	result,err:=db.Exec("update students set  name = ? where id = ?","gold" , 100)
	if err!=nil {
		log.Fatal(err)
		fmt.Println("修改失败")
	}
	//fmt.Println(result)
	fmt.Println(result.RowsAffected())
}
