package empty

import (
	"database/sql"
)
//引入mysql驱动程序
import _ "github.com/go-sql-driver/mysql"
/**
 * movie_db是数据库相关操作的功能文件
 */
/**
  数据库的操作：插入、查询、修改、删除
*/

//该函数核心作用：连接并打开数据库
func OpenDb()(*sql.DB,error){
	database, err :=sql.Open("mysql","root:123@tcp(127.0.0.1:3306)/movie?charset=utf8")
	if err != nil{
		//打开数据库遇到了错误
		return nil,err
	}
	return database,nil
}

//该函数的作用：将最新的电影的数据信息，更新到爬虫数据库当中，并返回结果
//布尔类型bool有两种可选值：true和false
//如果使用bool类型作为返回值类型，那么只能表示两种结果：成功、失败
//实际情况执行时，可能会出现多种情况，比如：
//   ① 数据有新的变化，此时update更新，会返回影响的行数代表真实数据发生改变的记录条数
//   ② 数据没有发生新的变化，此时update更新，sql语句执行能够成功，但是影响的行数却是0
//   ③ update语句执行时遇到问题, 影响的函数也是0
func UpdateMovieInfo(moveInfo MovieInfo)(int64,error){
	//1、打开数据库
	database,err := OpenDb()//openDb里面的所有数据，databases是起别名，放便理解
	if err != nil {
		return 0,err
	}
	//4、关闭数据库
	defer database.Close()

	//2、执行sql语句，更新数据库记录
	result, err :=database.Exec("update movie_info set movie_name = ?, movie_rate = ? where movie_id = ? ",
		moveInfo.Name,
		moveInfo.Rate,
		moveInfo.Id)
	//3、处理数据更新的结果
	if err != nil {
		return 0,err
	}
	rowNum, err := result.RowsAffected()
	if err != nil {
		return 0,err
	}
	return rowNum,nil
}

//根据电影编号查询数据电影数据，返回一个结构，布尔类型，true表示数据存在，false表示不存在
func QueryMovieInfo(id string) (bool,error){
	//1、打开数据库
	database,err := OpenDb()
	if err != nil{
		//打开数据库遇到了错误
		return false,err
	}
	//根据传递的电影的编号，执行数据库查询语句，判断数据中是否存在电影记录
	movieRow :=database.QueryRow("select movie_name from movie_info where movie_id = ? ",id)
	var movie_name string //默认值是空字符串

	err =movieRow.Scan(&movie_name)//如果movieRow当中没有数据，会产生错误：no rows in result set
	if movie_name != ""{//代表查询到了数据
		return true,nil
	}
	return false,err
}

//保存电影信息
func SaveMovieInfo(movie MovieInfo) (int64,error){
	//1、打开数据库
	database,err := OpenDb()
	if err != nil{
		//打开数据库遇到了错误
		return 0,err
	}

	//2、执行数据库操作
	result, err := database.Exec("insert into movie_info(" +
		"movie_id, movie_name, movie_rate) " +
		"values(?,?,?) ",movie.Id,movie.Name,movie.Rate)
	if err != nil{
		//遇到了错误
		return 0,err
	}
	//3、处理数据库执行结果
	id,err := result.RowsAffected()//执行sql语句，对数据操作影响了几行数据
	if err != nil{
		//遇到了错误
		return 0,err
	}

	//4、关闭数据库
	database.Close()
	return id,nil
}
