package main

import (
	"GolangApi/Controller"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	//Cors跨域設定
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
	})
	//handler實作
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	})

	//監聽Port設定
	http.HandleFunc("/CreateUser", Controller.CreateUser)
	http.ListenAndServe(":8778", c.Handler(handler))

}
