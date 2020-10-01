package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	coonstr:="root:123@tcp(127.0.0.1:3306)/school?charset=utf8"
	db,err:=sql.Open("mysql",coonstr)
	defer db.Close()
	if err!=nil {
		log.Fatal(err)
		fmt.Println("连接失败")
	}
	fmt.Println("连接成功")

   row:=db.QueryRow("select * from students  where id = ? ",100 )
   var id int
   var name string
   var sex  string
   var age  int
   var birthday  string
   var grade  float64
   err0:=row.Scan(&id,&name,&sex,&age,&birthday,&grade)
	if err!=nil {
		log.Fatal(err0)
		fmt.Println("查询失败")
		return
	}
	fmt.Println("学生编号:",id)
    fmt.Println("学生姓名:",name)
    fmt.Println("学生性别:",sex)
    fmt.Println("学生年龄:",age)
    fmt.Println("学生生日:",birthday)
    fmt.Println("学生成绩:",grade)
}