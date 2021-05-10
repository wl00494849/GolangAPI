package controller

import (
	"encoding/json"
	"fmt"
	"golang-api/server"
	"net/http"
)

var route = new(server.Route)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println("method:", r.Method)
	fmt.Println("path:", r.URL.Path)

	if r.Method == "GET" {
		route.RedirectRoute(w, "View/CreateUser.gtpl")
	} else {

		user := server.User{
			Account:  r.FormValue("Account"),
			Password: r.FormValue("Password"),
			Email:    r.FormValue("Email"),
		}

		user.CreateUser()

		//轉Json
		jsonData, _ := json.Marshal(user)
		fmt.Println("Json:", string(jsonData))

	}

	fmt.Println("code:", "200")
}

func UsersList(w http.ResponseWriter, r *http.Request) {

	userSlice := server.UsersList()

	jsonData, _ := json.Marshal(userSlice)

	w.Write(jsonData)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	var id int

	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&id)

	server.DeleteUser(id)

	fmt.Println("sucess")
}

func DockerTest(w http.ResponseWriter, r *http.Request) {

	jsonData, _ := json.Marshal("Hellow Docker")

	fmt.Println("Docker:sucess")

	w.Write(jsonData)
}
