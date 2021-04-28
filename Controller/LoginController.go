package Controller

import (
	"GolangApi/Server"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println("method:", r.Method)
	fmt.Println("path:", r.URL.Path)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("View/CreateUser.gtpl")
		log.Println(t.Execute(w, nil))
	} else {

		u := new(Server.Users)
		u.Account = r.FormValue("Account")
		u.Password = r.FormValue("Password")
		u.Email = r.FormValue("Email")

		u.CreateUser()

		//轉Json
		jsonData, _ := json.Marshal(u)

		w.Write(jsonData)
		fmt.Println("Json:", string(jsonData))
	}

	fmt.Fprintf(w, "200")
}
