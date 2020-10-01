package spider_db

import (
 "school/entity"
)

func QueryAdmin(name,pwd string) (*entity.Movie,error) {
	row :=MovieSpiderDb.QueryRow("select admin_name, admin_phone, admin_addr from movie_admin where admin_name = ? and admin_pwd = ?",name,pwd)
	var admin empty.Admin
	err := row.Scan(&admin.Name,&admin.Phone,&admin.Address)
	if err != nil {
		return nil,err
	}
	return &admin,nil

}
