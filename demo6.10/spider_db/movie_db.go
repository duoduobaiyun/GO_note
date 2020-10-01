package spider_db

import (
	"DoubanNewMovie/entity"
	"database/sql"
)
//导入驱动
import _ "github.com/go-sql-driver/mysql"

var MovieSpiderDb *sql.DB //全局的操作数据库的数据库对象
/**
 * 打开数据库
 */
func OpenDb()(error){
	if MovieSpiderDb != nil{//如果MovieSpiderDb不为nil，说明已经调用过OpenDB函数了，
		return nil
	}
	database,err :=sql.Open("mysql","root:yu123456@tcp(127.0.0.1:3306)/newmovie_class3?charset=utf8")
	if err != nil {
		return err
	}
	MovieSpiderDb = database
	return nil
}

//关闭数据库操作
func CloseDB()error{
	if MovieSpiderDb != nil{
		err :=MovieSpiderDb.Close()
		return err
	}
	return nil
}

/**
 * 查询数据库中电影的数据量
 */
func QueryMovieCount()(int,error){
	row:=MovieSpiderDb.QueryRow("select COUNT(movie.movie_id) count from movie")
	var count int
	err :=row.Scan(&count)
	if err != nil {
		return 0,err
	}
	return count,nil

}

/**
 * 保存电影信息到数据库
 */
func SavieMovie2Db(movie entity.Movie)(int64,error){
	////1、连接数据库
	//database,err :=OpenDb()
	//	if err != nil {
	//	return 0,err
	//}
	////4、关闭数据库
	//defer database.Close()

	//2、执行数据库的操作
	result, err := MovieSpiderDb.Exec("insert into movie" +
		"(movie_id, movie_name, movie_rate, movie_vote, movie_cover_url, movie_desc) " +
		"values" +
		"(?, ?, ?, ?, ?, ?)",
		movie.Id,movie.Name,movie.RateNum,movie.VoteNum,movie.CoverUrl,movie.Desc)
	//3、处理数据库操作的结果
	if err != nil {
		return 0,err
	}
	id, err :=result.RowsAffected()
	if err != nil {
		return 0,err
	}
	return id,nil

}
