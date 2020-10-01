package db_mysql

import (
	"fmt"
	"lol(new)/object1"
)

//根据账户和密码查询admin中数据
func QueryAdmin(userName,userPwd string)(*object1.Admin,error)  {
	row:=HeroDB.QueryRow("select admin_name, admin_phone, admin_username from admin where admin_userName = ? and admin_userPwd = ?",userName,userPwd)
	var admin object1.Admin
	err:=row.Scan(&admin.Name,&admin.UserName,&admin.Phone)
	if err!=nil {
		fmt.Println(err.Error())
		return nil,err
	}
	return &admin,nil
}

//根据账户和电话，email查询admin中数据
func QueryYZMAdmin(userName,phone,email string)(*object1.Admin,error) {
	row:=HeroDB.QueryRow("select admin_name,admin_username,admin_email from admin where admin_username=? and admin_phone=? and admin_email=? ",
		userName,phone,email)
	var admin object1.Admin
	err:=row.Scan(&admin.UserName,&admin.Phone,&admin.Email)
	if err!=nil {
		fmt.Println(err.Error())
		return nil,err
	}
	return &admin,nil
}

//查询所有用户
func QueryAdminCount()(int,error)  {
   row:=HeroDB.QueryRow("select count (admin.admin_id) count from admin")
	var count   int
   err:=row.Scan(&count)
	if err!=nil {
		fmt.Println(err.Error())
		return 0,err
	}
	return count,nil
}

func UpdateAdminInfo2DB(newPwd,userName string)  {
	HeroDB.Exec("update admin set admin_userPwd=? where admin_username=? ",newPwd,userName)
}

func SaveAdminInfo2DB(admin object1.Admin)(int64,error)  {
	result,err:=HeroDB.Exec("insert into admin " +
		"(admin_id,admin_name,admin_username,admin_userpwd,admin_phone,admin_sex,admin_email,admin_authority)" +
		"values" +
		"(?,?,?,?,?,?,?,?)",
		admin.Id,admin.Name,admin.UserName,admin.UserPwd,admin.Phone,admin.Sex,admin.Email,admin.Authority)
	if err!=nil {
		fmt.Println(err.Error())
		return 0,nil
	}
	id,err:=result.RowsAffected()
	if err!=nil {
		fmt.Println(err.Error())
		return 0,nil
	}
	return id,nil
}