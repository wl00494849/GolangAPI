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

	stmt, err := DataAcess.Db.Prepare("Insert Users set account=? ,password = ? ,email = ?")
	checkErr(err)

	stmt.Exec(u.Account, u.Password, u.Email)
	checkErr(err)

	DataAcess.Db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
