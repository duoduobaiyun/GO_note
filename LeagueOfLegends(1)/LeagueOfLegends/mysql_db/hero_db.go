package mysql_db

import (
	"LeagueOfLegends/object"
	"database/sql"
	"fmt"
	_ "go-sql-driver/mysql"
)

var HeroDb *sql.DB//全局数据库操作对象

//1.打开数据库
func SqlOpen () error {
	if HeroDb != nil {
		return nil
	}
	detabase,err := sql.Open("mysql","root:zjt3606810515@tcp(127.0.0.1:3306)/hero?charset=utf8")
	if err != nil {
		return err
	}
	HeroDb = detabase
	return nil
}

//2.关闭数据库
func SqlClose() error {
	if HeroDb != nil {
		err := HeroDb.Close()
		return err
	}
	return nil
}

//3.存储数据信息
func SavieHeroInfo2Db (hero object.Hero) (int64,error) {
	result, err := HeroDb.Exec("insert  into heroinfo" +
		"(heroId,name,alias,title,attack,defense,magic,difficulty,selectAudio,banAudio)" +
		"values" +
		"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		hero.HeroId,hero.Name,hero.Alias,hero.Title,hero.Attack,hero.Defense,hero.Magic,
		hero.Difficulty,hero.SelectAudio,hero.BanAudio)
	if err != nil {
		fmt.Println(err.Error())
		return 0,err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0,err
	}
	return id,nil
}

//4.查询数据库中hero的数量
func QueryHeroCount () (int,error) {
	row := HeroDb.QueryRow("select COUNT(heroinfo.heroId) count from heroinfo")
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0,err
	}
	return count,nil
}

//5.查询数据库中的数据
func QueryAllHeros () ([]object.Hero,error) {
	heroSlice := make([]object.Hero,0)
	rows, err := HeroDb.Query("select * from heroinfo")
	if err != nil {
		return heroSlice, nil
	}
	for rows.Next() {
		var hero object.Hero
		err = rows.Scan(&hero.HeroId,&hero.Name,&hero.Alias,&hero.Title,&hero.Attack,&hero.Defense,&hero.Magic,&hero.Difficulty,
			&hero.SelectAudio,&hero.BanAudio)
		if err != nil {
			return heroSlice, nil
		}
		heroSlice = append(heroSlice,hero)
	}
	return heroSlice, nil
}

