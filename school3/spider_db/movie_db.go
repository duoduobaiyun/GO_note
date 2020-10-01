package spider_db

import (
	"database/sql"
	"school3/entity"
)
import _"github.com/go-sql-driver/mysql"

var MovieSpiderDb *sql.DB


func OpenDb() error {
	database,err :=sql.Open("mysql","root:123@tcp(127.0.0.1:3306)/new_table?charset=utf8")
	if err != nil {
		return err
	}
	MovieSpiderDb = database
	return nil
}

func CloseDB() error {
	if MovieSpiderDb != nil {
		err := MovieSpiderDb.Close()
		return err
	}
	return nil
}

func QueryMovieCount() (int,error) {
	row:= MovieSpiderDb.QueryRow("select COUNT(new_table.movie_id) count_num from new_table")
	var count int
	err := row.Scan(&count)
	if err!= nil {
		return 0,err
	}
	return count,nil
}
//
func QueryAllMovies()([]entity.Movie,error){
	movieSlice := make([]entity.Movie,0)
	rows,err :=MovieSpiderDb.Query("select * from new_table")
	if err != nil {
		return movieSlice,err
	}
	//数组、切片
	for rows.Next(){//bool
		var movie entity.Movie
		err =rows.Scan(&movie.Id,&movie.Name,&movie.RateNum,&movie.VoteNum,&movie.CoverUrl,&movie.Desc)
		if err != nil {
			return movieSlice,err
		}
		//把movie数据追加到切片中
		movieSlice = append(movieSlice,movie)
	}
	return movieSlice,nil
}
//


func SavidMovie2Db(movie entity.Movie) (int64 , error) {
	result , err := MovieSpiderDb.Exec("insert into new_table" +
		"(movie_id,movie_name,movie_rate,movie_vote,movie_cover,movie_desc) "+
		"values"+
	  "(?,?,?,?,?,?)",
	  movie.Id,movie.Name,movie.RateNum,movie.VoteNum,movie.CoverUrl,movie.Desc)
	if err!= nil {
		return 0 ,err
	}
	id,err:=result.RowsAffected()
	if err!=nil {
		return 0,nil
	}
	return id,nil
}