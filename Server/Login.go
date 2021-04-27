package Server

import (
	"GolangApi/DataAcess"
)

type Users struct {
	Account  string
	Password string
	Email    string
}

func (u *Users) CreateUser() {

	conn := DataAcess.Db

	stmt, err := conn.Prepare("Insert Users set account=? ,password = ? ,email = ?")
	checkErr(err)

	stmt.Exec(u.Account, u.Password, u.Email)
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
