package spider_lol

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"project1/LoL/object"
)

var SpiderDb *sql.DB

func OpenDB()error {
	if SpiderDb !=nil{
		return nil
	}
	database,err:=sql.Open("mysql","root:123@tcp(127.0.0.1:3306)/lol?charset=utf8")
	if err != nil {
		return err
	}
	SpiderDb =database
	return nil
}

func CloseDB() error {
	if SpiderDb !=nil{
		err:=SpiderDb.Close()
		return err
	}
	return nil
}

func QueryAllHero()([]object.Loldatabase,error) {
	heroSlice := make([]object.Loldatabase, 0)
	rows, err :=SpiderDb.Query("select * from lolhero")
	if err!=nil {
		return  heroSlice,err
	}
	for rows.Next() {
		var hero object.Loldatabase
		err=rows.Scan(&hero.Id,&hero.Name,&hero.Alias,&hero.Title,&hero.VoiceOne,&hero.VoiceTwo,&hero.Label,&hero.ImgUrl)
		if err!=nil{
			return  heroSlice,err
		}
		heroSlice=append(heroSlice,hero)
	}
	return heroSlice,nil

}

func QueryHeroCount() (int, error) {
	row:=SpiderDb.QueryRow("select COUNT(lolhero.lol_id) count from  lolhero")
	var count int
	err:=row.Scan(&count)
	if err != nil {
		return 0,err
	}
	return  count,nil
}


func SaveHero(hero object.Loldatabase) (int64,error) {
	result,err :=SpiderDb.Exec("insert into lolhero"+
		"(lol_id,lol_name,lol_alias,lol_title,lol_voiceOne,lol_voiceTwo,lol_label,imgurl)"+
		"values"+
		"(?,?,?,?,?,?,?,?)",
		hero.Id,hero.Name,hero.Alias,hero.Title,hero.VoiceOne,hero.VoiceTwo,hero.Label,hero.ImgUrl)
	if err!=nil{
		return  0,nil
	}
	id,err:=result.RowsAffected()
	if err != nil {
		return 0,err
	}
	return id,nil
}