package main

import (
	"GolangApi/Controller"
	"GolangApi/Server"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

var er = new(Server.ErrorHandle)

type Mux struct{}

func main() {
	//Cors跨域設定
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200", "http://localhost:8778"},
	})

	mux := &Mux{}

	//監聽Port設定
	err := http.ListenAndServe(":8778", c.Handler(mux))
	er.CheckErr(err)

	http.HandleFunc("/CreateUser", Controller.CreateUser)
	fmt.Println("Port:", "Ok")

}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/CreateUser":
		Controller.CreateUser(w, r)
		break
	default:
		http.NotFound(w, r)
	}

}
