package main

import (
	"GolangApi/Controller"
	"net/http"
)

func main() {

	http.HandleFunc("/CreateUser", Controller.CreateUser)
	http.ListenAndServe(":8778", nil)

}
