package Controller

import (
	"GolangApi/Server"
	"encoding/json"
	"fmt"
	"net/http"
)

var route = new(Server.Route)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println("method:", r.Method)
	fmt.Println("path:", r.URL.Path)

	if r.Method == "GET" {
		route.RedirectRoute(w, "View/CreateUser.gtpl")
	} else {

		u := Server.User{
			Account:  r.FormValue("Account"),
			Password: r.FormValue("Password"),
			Email:    r.FormValue("Email"),
		}

		u.CreateUser()

		//轉Json
		jsonData, _ := json.Marshal(u)
		fmt.Println("Json:", string(jsonData))

	}

	fmt.Println("code:", "200")
}

func UsersList(w http.ResponseWriter, r *http.Request) {

	userSlice := Server.UsersList()

	jsonData, _ := json.Marshal(userSlice)

	w.Write(jsonData)
}
