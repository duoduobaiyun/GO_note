package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main(){
  connstr:="root:123@tcp(127.0.0.1:3306)/college?charset=utf8"
  db,err:=sql.Open("mysql",connstr)
   defer db.Close()
	if err!=nil  {
		fmt.Println("连接失败")
		log.Fatal(err)
		return
	}
	fmt.Println("连接成功")

    insert,err:=db.Prepare("insert into grade(id, stu_name, stu__class, grade, db_mysql)values (3,'liu','chinese',70,180)")
	if err!=nil {
		log.Fatal(err)
		fmt.Println("插入数据执行失败")
		return
	}

	result,err:=insert.Exec()
	if err!=nil {
		log.Fatal(err)
		fmt.Println("插入数据执行失败")
		return
	}

	rowAffected,err:=result.RowsAffected()
	if err!=nil  {
		log.Fatal(err)
	}
	fmt.Printf("影响了%d行\n",rowAffected)



}
