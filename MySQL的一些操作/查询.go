package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	conner:="root:123@tcp(127.0.0.1:3306)/wish?charset=utf8"
	db,err:=sql.Open("mysql",conner)
	if err!=nil {
		log.Fatal(err)
		fmt.Println("连接失败")
	}

	result:=db.QueryRow("select * from student where id =?",600)
	var id int
	var name string
	var sex  string
	var grade  int
	err1 := result.Scan(&id,&name,&sex,&grade)
	if err1!=nil {
		log.Fatal(err)
		fmt.Println("查询失败")
	}
	fmt.Println("学生编号:",id)
	fmt.Println("学生性别",sex)
	fmt.Println("学生姓名:",name)
	fmt.Println("学生成绩:",grade)
}
