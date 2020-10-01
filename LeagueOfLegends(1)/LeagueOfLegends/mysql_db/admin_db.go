package mysql_db

import (
	"LeagueOfLegends/object"
	"fmt"

	//"database/sql"
	_ "go-sql-driver/mysql"
)


//根据账户和密码查询admin中数据
func QueryAdmin (userName,userPwd string) (*object.Admin,error) {
	row := HeroDb.QueryRow("select admin_name, admin_phone, admin_userName from admin where admin_userName = ? and admin_userPwd = ?",userName,userPwd)

	var admin object.Admin
	err := row.Scan(&admin.Name,&admin.UserName,&admin.Phone)
	if err != nil {
		return nil,err
	}
	return &admin,nil
}

//根据账户和电话，email查询admin中数据
func QueryYZMAdmin (userName,phone,email string) (*object.Admin,error) {
	row := HeroDb.QueryRow("select admin_name,admin_userName,admin_email from admin where admin_userName = ? and admin_phone = ? and admin_email = ?",
		userName,phone,email)

	var admin object.Admin
	err := row.Scan(&admin.Name,&admin.UserName,&admin.Email)
	if err != nil {
		return nil,err
	}
	return &admin,nil
}

func QueryAdminCount () (int,error) {
	row := HeroDb.QueryRow("select COUNT(admin.admin_id) count from admin")
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0,err
	}
	return count,nil
}

func UpdateAdminInfo2Db (newPwd,userName string){
	HeroDb.Exec("update admin set admin_userPwd = ? where admin_userName = ?",newPwd,userName)
}

func SavieAdminInfo2Db (admin object.Admin) (int64,error) {
	result, err := HeroDb.Exec("insert  into admin" +
		"(admin_id,admin_name,admin_userName,admin_userPwd,admin_phone,admin_sex,admin_email,authority)" +
		"values" +
		"(?,?,?,?,?,?,?,?)",
		admin.Id,admin.Name,admin.UserName,admin.UserPwd,admin.Phone,admin.Sex,admin.Email,admin.Authority)
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