package mysql_db

import (
	"LOLproject/object"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var LOL *sql.DB

func Open_db() error {
	database, err := sql.Open("mysql", "root:wenb1125@tcp(127.0.0.1:3306)/lol_info?charset=utf8")
	if err != nil {
		return nil
	}
	LOL = database
	return nil
}
func Close_db() error {
	if LOL != nil {
		err := LOL.Close()
		return err
	}
	return nil
}
func SelectAll() ([]object.Hero, error) {
	Herolol := make([]object.Hero, 0)
	rows, err := LOL.Query("select * from lol_info")
	if err != nil {
		return Herolol, err
	}
	var hero object.Hero
	for rows.Next() {

		err = rows.Scan(&hero.HeroId, &hero.Name, &hero.Alias, &hero.Title, &hero.IsWeekFree, &hero.SelectAudio, &hero.BanAudio)
		if err != nil {
			fmt.Println(err.Error())
			return Herolol, err
		}
		Herolol = append(Herolol, hero)
	}
	return Herolol, nil
}
func SelectRows()(int,error)  {
	var count int
	rows := LOL.QueryRow("select COUNT(lol_info.hero_id) from lol_info")
	err := rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
func Hero_db(hero object.Hero) (int64, error) {
	rezult, err := LOL.Exec("insert into lol_info" +
		"(hero_id, hero_big_name, hero_english_name, hero_small_name, hero_isweekfree, hero_voice_1, hero_voice_2)"+
		"values" +
		"(?, ?, ?, ?, ?, ?, ?)",
		hero.HeroId, hero.Name, hero.Alias, hero.Title, hero.IsWeekFree, hero.SelectAudio, hero.BanAudio)
	if err != nil {
		fmt.Println(err.Error())
		return 0, nil
	}
	id, err := rezult.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return id,nil
}
