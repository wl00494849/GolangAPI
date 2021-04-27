package Server

import (
	"database/sql"
)

type Users struct {
	Account  string
	Password string
	Email    string
}

func (u *Users) CreateUser() {

	db, err := sql.Open("mysql", "root:1qaz2wsx3edc@/projectmanage?charset=utf8")
	checkErr(err)

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
