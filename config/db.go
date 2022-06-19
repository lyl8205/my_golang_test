package config

import (
	"os"

	mysql2 "codeup.aliyun.com/5f69c1766207a1a8b17fda8e/sanhe_library/mysql"
)

type mysql struct {
	ShanHeErpFlow, SanHeGdMobile *mysql2.MysqlDialect
}

var Mysql = &mysql{
	ShanHeErpFlow: &mysql2.MysqlDialect{
		Host: os.Getenv("mysql_host"),
		User: os.Getenv("mysql_username"),
		Pwd:  os.Getenv("mysql_password"),
		Db:   os.Getenv("mysql_erp_flow_database"),
		Port: os.Getenv("mysql_port"),
	},
	SanHeGdMobile: &mysql2.MysqlDialect{
		Host: os.Getenv("mysql_host"),
		User: os.Getenv("mysql_username"),
		Pwd:  os.Getenv("mysql_password"),
		Db:   os.Getenv("mysql_gdmobile_database"),
		Port: os.Getenv("mysql_port"),
	},
}
