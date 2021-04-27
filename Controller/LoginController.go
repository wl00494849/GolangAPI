package Controller

import (
	"GolangApi/Server"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func CreateUserController(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println("method:", r.Method)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("View/CreateUser.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		u := new(Server.Users)
		u.Account = r.FormValue("Account")
		u.Password = r.FormValue("Password")
		u.Email = r.FormValue("Email")

		u.CreateUser()
	}

	fmt.Fprintf(w, "200")
}
