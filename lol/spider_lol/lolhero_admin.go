package spider_lol

import "project1/LoL/object"

func QueryAdmin(name, pwd string) (*object.Admin, error) {
	row:=SpiderDb.QueryRow("select admin_name, admin_phone, admin_addr from lolhero_admin where admin_name = ? and admin_pwd = ?",name,pwd)
	var admin object.Admin
	err :=row.Scan(&admin.Name,&admin.Phone,&admin.Address)
	if err !=nil{
		return nil,err
	}
	return &admin,nil
}
