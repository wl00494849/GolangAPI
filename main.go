package main

import (
	"GolangApi/Controller"
	"GolangApi/Server"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

var er = new(Server.ErrorHandle)

func main() {
	//Cors跨域設定
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200", "http://localhost:8778"},
	})
	fmt.Println("Cors:", "Ok")

	//handler實作
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	})
	fmt.Println("handle:", "Ok")

	//監聽Port設定
	http.HandleFunc("/CreateUser", Controller.CreateUser)
	fmt.Println("Port:", "Ok")

	err := http.ListenAndServe(":8778", c.Handler(handler))
	er.CheckErr(err)

}
