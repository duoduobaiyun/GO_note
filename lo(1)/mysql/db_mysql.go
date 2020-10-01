package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"lo(1)/object"
)

var HeroDB *sql.DB //全局数据库操作对象

//打开数据库
func OPenDB() error {
	databases, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/lol(1)?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("打开数据库")
	HeroDB = databases
	return nil
}

//关闭数据库
func CloseDB() error {
	if HeroDB != nil {
		err := HeroDB.Close()
		fmt.Println("关闭数据库")
		return err
	}
	return nil
}

//查询所有英雄信息
func QueryAllHero() ([]object.Name, error) {
	heroSlice := make([]object.Name, 0)
	rows, err := HeroDB.Query("select * from heroinfo")
	if err != nil {
		fmt.Println(err.Error())
		return heroSlice, err
	}
	for rows.Next() {
		var hero object.Name
		err = rows.Scan(&hero.Id, &hero.ImgUrl, &hero.Name, &hero.Alias, &hero.Title, &hero.Attack, &hero.Magic)
		if err != nil {
			return heroSlice, err
		}
		heroSlice = append(heroSlice, hero)
	}
	return heroSlice, nil
}

//查询所有的英雄的数量
func QueryHeroCount() (int, error) {
	row := HeroDB.QueryRow("select COUNT (heroinfo.id) count from heroinfo")
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
