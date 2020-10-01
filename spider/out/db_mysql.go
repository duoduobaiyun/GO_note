package out

import (
	"database/sql"
	"log"
	_"github.com/go-sql-driver/mysql"
)
/**
 * movie_db是数据库相关操作的功能文件
 */
/**
  数据库的操作：插入、查询、修改、删除
*/

//该函数核心作用：连接并打开数据库
func OpenDB()(*sql.DB,error)  {
	databases,err:=sql.Open("mysql","root:123@tcp(127.0.0.1:3306)/movie?charset=utf8")
	if err!=nil {
		log.Fatal(err)
		//打开数据库遇到了错误
		return nil,err
	}
	return databases,nil
}


//该函数的作用：将最新的电影的数据信息，更新到爬虫数据库当中，并返回结果
//布尔类型bool有两种可选值：true和false
//如果使用bool类型作为返回值类型，那么只能表示两种结果：成功、失败
//实际情况执行时，可能会出现多种情况，比如：
//   ① 数据有新的变化，此时update更新，会返回影响的行数代表真实数据发生改变的记录条数
//   ② 数据没有发生新的变化，此时update更新，sql语句执行能够成功，但是影响的行数却是0
//   ③ update语句执行时遇到问题, 影响的函数也是0

func UpdateMovieInfo(movieInfo MovieInfo)(int64,error)  {
	//1、打开数据库
      databases,err:=OpenDB()
	  if err!=nil {
		log.Fatal(err)
		return 0,err
	}
	//4、关闭数据库
	defer databases.Close()
	//2、执行sql语句，更新数据库记录
	result,err:=databases.Exec("update movie set movie_name=?,movie_ratReg=?,movie_vote=?,movie_desc=?,movie_image=? where movie_id=?",
		movieInfo.Name,
		movieInfo.RatReg,
		movieInfo.Vote,
		movieInfo.Desc,
		movieInfo.Image,
		movieInfo.Id)
	if err!=nil {
		log.Fatal(err)
		return 0,err
	}
	rowNum,err:=result.RowsAffected()
	if err!=nil {
		log.Fatal(err)
		return 0,err
	}
	return rowNum,nil
}
//根据电影编号查询数据电影数据，返回一个结构，布尔类型，true表示数据存在，false表示不存在
func QueryMovieInfo(id string)(bool,error){
     //打开数据库
	databases,err:=OpenDB()
	if err!=nil {
		log.Fatal(err)
		//打开数据库遇到了错误
		return false,err
	}
	//根据传递的电影的编号，执行数据库查询语句，判断数据中是否存在电影记录
	movieRow:=databases.QueryRow("select  movie_name from movie  where  movie_id=?",id)
	var  movie_name string //默认值是空字符串

	err=movieRow.Scan(&movie_name)//如果movieRow当中没有数据，会产生错误：no rows in result set
	if movie_name!="" {
		return true,nil
	}
	return false,err
}

func SaveMovieInfo(movie  MovieInfo)(int64,error)  {
	//1、打开数据库
	databases,err:=OpenDB()
	if err!=nil {
		log.Fatal(err)
		return 0,err
	}
	result,err:=databases.Exec("insert into movie ("+"movie_id,movie_name,movie_ratReg,movie_vote,movie_desc,movie_image)"+
		"values(?,?,?,?,?,?)",movie.Id,movie.Name,movie.RatReg,movie.Vote,movie.Desc,movie.Image)
	if err!=nil {
		log.Fatal(err)
		//遇到了错误
		return 0,err
	}
	//3、处理数据库执行结果
	rowNum,err:=result.RowsAffected()
	if err!=nil {
		log.Fatal(err)
		//遇到了错误
		return 0,err
	}
	defer databases.Close()
	return rowNum,nil
}