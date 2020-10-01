package spider_db

import (
	"school3/entity"
)

func QueryAdmin(name,pwd string) (*entity.Admin,error) {
	row :=MovieSpiderDb.QueryRow("select admin_name, admin_phone, admin_addr from movie_admin where admin_name = ? and admin_pwd = ?",name,pwd)
	var admin entity.Admin
	err := row.Scan(&admin.Name,&admin.Phone,&admin.Address)
	//fmt.Println(pwd,name)
	if err != nil {
		return nil,err
	}
	//fmt.Println(admin)
	return &admin,nil

}
