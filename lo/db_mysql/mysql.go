package db_mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"project1/lo/object"
)

var LolDB *sql.DB

func OPenDB()error  {
   databases,err:=sql.Open("mysql","root:123@tcp(127.0.0.1:3306)/lol?charset=utf8")
	if err!=nil {
		fmt.Println(err.Error())
		return err
	}
	LolDB=databases
	return nil
}

func CloseDB()error  {
	if LolDB!=nil {
     err:=LolDB.Close()
     return err

	}
	return nil
}

func QueryAllHero()([]object.Names,error) {
	heroSlice := make([]object.Names, 0)
	rows, err :=LolDB.Query("select * from Names ")
	if err!=nil {
		return  heroSlice,err
	}
	for rows.Next() {
		var hero object.Names
		err=rows.Scan(&hero.Id,&hero.ImgUrl,&hero.Name,&hero.Alias,&hero.Title,&hero.Attack,&hero.Defense,&hero.Magic)
		if err!=nil{
			return  heroSlice,err
		}
		heroSlice=append(heroSlice,hero)
	}
	return heroSlice,nil

}

func QueryHeroCount() (int, error) {
	row:=LolDB.QueryRow("select COUNT(Names .id) count from  Names")
	var count int
	err:=row.Scan(&count)
	if err != nil {
		return 0,err
	}
	return  count,nil
}

func SaveHero(hero object.Names) (int64,error) {
	result,err :=LolDB.Exec("insert into names"+
		"(ImgUrl,Name,Alias,Title,Attack,Defense,Magic) "+
		"values"+
		"(?,?,?,?,?,?,?)",
		hero.ImgUrl,hero.Name,hero.Alias,hero.Title,hero.Attack,hero.Defense,hero.Magic)
	if err!=nil{
		fmt.Println(err.Error())
		return  0,err
	}
	id,err:=result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return 0,err
	}
	fmt.Println("保存英雄:",id)
	return id,nil
}