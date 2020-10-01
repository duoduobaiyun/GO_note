package spider_db

import (
	"database/sql"
	_struct "demo6.12/struct"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var MovieDb*sql.DB //全局操控的数据库
//打开数据库
func OpenDb()error  {
	if MovieDb!=nil {
		return nil//不为nil的时候OpenDb是已经被调用
	}
	databases,err:=sql.Open("mysql","root:123@tcp(127.0.0.1:3306)/movie?charset=utf8")
	if err!=nil {
		log.Fatal(err)
		return nil
	}
	MovieDb=databases
	return nil
}

//关闭数据库
func CloseDb()error  {
	if MovieDb!=nil {
		err:=MovieDb.Close()
		return err
	}
	return nil
}

func QueryMovieCount()(int,error)  {
	row:=MovieDb.QueryRow("select COUNT (movie.movie_id) count_num from movie")
	var count  int//初始值
	err:=row.Scan(&count)
	if err!=nil {
		return 0,err
	}
	return count,nil
}

func SaveDB(movie _struct.Movie)(int64,error)  {
	result,err:=MovieDb.Exec("inset into  movie (movie_id , movie_name, movie_ratReg,movie_vote,movie_desc,movie_image) " +
		"values" +
		"(?,?,?,?,?,? )",
		movie.Id,movie.Name,movie.RatReg,movie.Vote,movie.Desc,movie.Image)
	//处理结果
	if err!=nil {
		return 0,err
	}
	id,err:=result.RowsAffected()
	if err!=nil {
		log.Fatal(err)
		return 0,err
	}
    return id,nil
}