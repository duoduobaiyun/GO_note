package db_mysql

import (
	"database/sql"
	"log"
)

func OpenDB()(*sql.DB,error)  {
	databases,err:=sql.Open("mysql","root:123@tcp(127.0.0.1:3306)/movie?charset=utf8")
	if err!=nil {
		log.Fatal(err)
		return nil,err
	}
	return databases,nil
}

