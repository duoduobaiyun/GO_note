package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/**
 *  定义：学生结构类型
 *  封装：
 */
type Student struct{
	id int
	name string
	sex string
	grade  int
}

func main() {

	connStr := "root:123@tcp(127.0.0.1:3306)/wish?charset=utf8"
	db, err :=sql.Open("mysql",connStr)
	if err != nil{
		fmt.Println("数据库连接失败：",err.Error())
		return
	}

	/**
	  目标：查询出stu_info中所有的数据记录
	*/
	//没有？时，可变参数类型可以不填
	rows, err :=db.Query("select * from student")
	if err != nil {
		fmt.Println("数据库查询失败：",err.Error())
		return
	}
	//查询数据成功，得到了数据
	//column:列
	//查询到的数据中包含的字段名，
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println("获取失败：",err.Error())
		return
	}
	fmt.Println(columns)


	dataSlice := make([]Student,0)

	//循环
	for rows.Next(){//当Next为true的时候，表示条件满足，执行for循环中的代码
		var id int
		var name string
		var sex string
		var grade  int
		err := rows.Scan(&id,&name,&sex,&grade)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		student := Student{id,name,sex,grade}
		dataSlice = append(dataSlice,student)
		//读取成功
		//fmt.Println(id,name,sex,birthday)
	}

	//打印slice
	fmt.Println(dataSlice)
	for _,student := range dataSlice{
		fmt.Println(student.id,student.name,student.sex,student.grade)
	}

}
