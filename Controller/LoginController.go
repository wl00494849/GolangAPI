package Controller

import (
	"GolangAPI/Server"
	"fmt"
	"net/http"
)

func CreateUserController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	u := new(Server.Users)
	u.Account = r.FormValue("Account")
	u.Password = r.FormValue("Password")
	u.Email = r.FormValue("Email")

	u.CreateUser()

	fmt.Fprintf(w, "200")
}
