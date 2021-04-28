package main

import (
	"GolangApi/Controller"
	"net/http"

	"github.com/rs/cors"
)

func main() {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
	})

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	})

	http.HandleFunc("/CreateUser", Controller.CreateUser)
	http.ListenAndServe(":8778", c.Handler(handler))

}
