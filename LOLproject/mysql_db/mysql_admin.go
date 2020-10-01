package mysql_db

import (
	"LOLproject/object"
)

func Admin(name, pwd string) (*object.Admin, error) {
	row := LOL.QueryRow("select admin_name, admin_phone, admin_address from admin where admin_name = ? and admin_pwd = ?", name, pwd)

	var admin object.Admin
	err := row.Scan(&admin.Username, &admin.Phone, &admin.Address)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
