package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	connstr := "root:123@tcp(127.0.0.1:3306)/wish?charset=utf8"
	db, err := sql.Open("mysql", connstr)
	if err != nil {
		log.Fatal(err)
		fmt.Println("连接失败")
	}

	result, err := db.Query("select * from student")
	if err != nil {
		log.Fatal(err)
		fmt.Println("查询失败")
	}
	//只查询到字段名
	columns, err := result.Columns()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(columns)

	isExit := result.Next()
	if isExit {
		var id int
		var name string
		var sex string
		var grade int
		result.Scan(&id, &name, &sex, &grade)
		fmt.Println(id, name, sex, grade)

		isExit = result.Next()
		if isExit { //第二条数据
			id = 0
			name = ""
			sex = ""
			grade=0
			result.Scan(&id, &name, &sex,&grade)
			fmt.Println(id, name, sex,grade)
		}
	}
}