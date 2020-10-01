package db_mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"lol(new)/object1"
)

var HeroDB *sql.DB//全局数据库操作对象

//打开数据库
func OpenDB() error {
	databases, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/lol(new)?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	HeroDB = databases
	return nil
}

//关闭数据库
func CloseDB() error {
	if HeroDB != nil {
		err := HeroDB.Close()
		return err
	}
	return nil
}

//查询所有英雄的数量
func QueryHeroCount() (int, error) {
	row := HeroDB.QueryRow("select count(new_lol.heroid) count from new_lol")
	var count int
	err := row.Scan(&count)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return 0, nil
}

//存储数据英雄数据
func SaveHeroInfo2DB(hero object1.Hero) (int64, error) {
	result, err := HeroDB.Exec("insert into new_lol "+
		"(heroId,name,alias,title,attack,defense,magic,difficulty,selectAudio,banAudio)"+
		"values"+
		"(?,?,?,?,?,?,?,?,?,?)",
		hero.HeroId, hero.Name, hero.Alias, hero.Title, hero.Attack, hero.Defense, hero.Magic, hero.Difficulty, hero.SelectAudio, hero.BanAudio)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return id, nil
}

//查询数据库中所有英雄的数据
func QueryAllHero() ([]object1.Hero, error) {
	heroSlice := make([]object1.Hero, 0)//容器
	row, err := HeroDB.Query("select from new_lol")
	if err != nil {
		fmt.Println(err.Error())
		return heroSlice, nil
	}
	for row.Next() {
		var hero object1.Hero//数据
		err = row.Scan(&hero.HeroId, &hero.Name, &hero.Alias, &hero.Title, &hero.Attack, &hero.Defense, &hero.Magic, &hero.Difficulty, &hero.SelectAudio, &hero.BanAudio)
		if err!=nil {
			fmt.Println(err.Error())
			return heroSlice,nil
		}
		heroSlice=append(heroSlice,hero)//数据存放在容器里面
	}
	return heroSlice,nil
}
