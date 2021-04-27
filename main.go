package main

import (
	"GolangApi/Controller"
	"net/http"
)

func main() {

	http.HandleFunc("/CreateUser", Controller.CreateUserController)
	http.ListenAndServe(":8778", nil)

}
