package spider_db

import (
	_struct "demo6.12/struct"
)

func QueryAdmin(name,pwd string)(*_struct.Admin,error){
	//1、执行sql语句
	row :=MovieDb.QueryRow("select admin_name, admin_phone,admin_pwd  from admin where admin_name = ? and admin_pwd = ?",name,pwd)

	var admin _struct.Admin
	err := row.Scan(&admin.Name,&admin.Pwd,&admin.Id)
	if err != nil {
		return nil,err
	}
	return &admin,nil
}
