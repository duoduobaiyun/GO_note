package util

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var LolDB *sql.DB

/**
 * 打开数据库
 */
func OpenDb() error {
	if LolDB != nil {
		return nil
	}
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/lol_info?charset=utf8")
	if err != nil {
		return err
	}
	LolDB = db
	return nil
}

/**
 * 关闭数据库
 */
func CloseDb() error {
	if LolDB != nil {
		err := LolDB.Close()
		if err != nil {
			return err
		}
		LolDB = nil
	}
	return nil
}
