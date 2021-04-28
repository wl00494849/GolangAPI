package Server

import (
	"GolangApi/DataAcess"
)

type Users struct {
	Account  string
	Password string
	Email    string
}

var er = new(ErrorHandle)

func (u *Users) CreateUser() {

	conn := DataAcess.Db

	stmt, err := conn.Prepare("Insert Users set account=? ,password = ? ,email = ?")
	er.CheckErr(err)

	stmt.Exec(u.Account, u.Password, u.Email)
	er.CheckErr(err)

}
