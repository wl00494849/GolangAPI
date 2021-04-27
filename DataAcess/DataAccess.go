package DataAcess

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func GetDbConn() *sql.DB {
	config := mysql.Config{
		User:                 "root",
		Passwd:               "1qaz2wsx3edc",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "projectmanage",
		AllowNativePasswords: true,
	}

	db, _ := sql.Open("mysql", config.FormatDSN())
	return db
}
