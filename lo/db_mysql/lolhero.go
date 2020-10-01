package db_mysql

import (
	"fmt"
	"project1/lo/object"
)

func QueryAdmin(name, pwd string) (*object.Admin, error) {
	row:=LolDB.QueryRow("select admin_name, admin_phone, admin_address from Admin where admin_name = ? and admin_pwd = ?",name,pwd)
	var admin object.Admin
	fmt.Println(admin)
	err :=row.Scan(&admin.Name,&admin.Pwd,&admin.Address)
	if err !=nil{
		fmt.Println(err.Error())
		return nil,err
	}
	return &admin,nil
}

//func QueryYZMAdmin(username,phone,email string)(*object.Admin,error){
//	row := mysql_db.HeroDb.QueryRow("select admin_name,admin_userName,admin_email from admin where admin_userName = ? and admin_phone = ? and admin_email = ?",
//		userName,phone,email)
//}