package Server

import (
	"GolangApi/DataAccess"
	"GolangApi/Model"
)

type User struct {
	Account  string
	Password string
	Email    string
}

func (u *User) CreateUser() {

	conn := DataAccess.Db

	stmt, err := conn.Prepare("Insert Users set account=? ,password = ? ,email = ?")
	er.CheckErr(err)

	stmt.Exec(u.Account, u.Password, u.Email)
	er.CheckErr(err)

}

func UsersList() []Model.User {
	userSlice := make([]Model.User, 0)
	conn := DataAccess.Db

	row, err := conn.Query("SELECT * FROM users")
	er.CheckErr(err)

	for row.Next() {
		user := new(Model.User)

		ScanToStruct(row, user)

		userSlice = append(userSlice, *user)
	}

	return userSlice
}
