package mysql

func QueryAdmin(name,pwd string)  {
	row:=HeroDB.QueryRow("select admin_id,admin_name,admin_pwd ")
}
