package dataAccess

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
		ParseTime:            true,
	}

	db, _ := sql.Open("mysql", config.FormatDSN())

	//最大開幾個連線
	db.SetMaxOpenConns(5)
	//最多幾個閒置連線
	db.SetMaxIdleConns(2)
	//閒置多久後刪除連線
	db.SetConnMaxLifetime(time.Hour)

	Db = db
}
