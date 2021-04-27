package DataAcess

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	config := mysql.Config{
		User:                 "root",
		Passwd:               "1qaz2wsx3edc",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "projectmanage",
		AllowNativePasswords: true,
	}

	db, _ := sql.Open("mysql", config.FormatDSN())

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(time.Hour)

	Db = db
}
