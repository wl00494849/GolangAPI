package Server

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type Users struct {
	Account  string
	Password string
	Email    string
}

func (u *Users) CreateUser() {

	config := mysql.Config{
		User:                 "root",
		Passwd:               "1qaz2wsx3edc",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "projectmanage",
		AllowNativePasswords: true,
	}

	db, _ := sql.Open("mysql", config.FormatDSN())

	stmt, err := db.Prepare("Insert Users set account=? ,password = ? ,email = ?")
	checkErr(err)

	stmt.Exec(u.Account, u.Password, u.Email)
	checkErr(err)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
