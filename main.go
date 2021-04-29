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
		AllowedOrigins: []string{"http://localhost:4200"},
	})
	fmt.Println("Cors:", "Ok")

	mux := &Mux{}

	//監聽Port設定
	err := http.ListenAndServe(":8778", c.Handler(mux))
	er.CheckErr(err)

}

//路由
func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/CreateUser":
		Controller.CreateUser(w, r)
		break
	case "/UsersList":
		Controller.UsersList(w, r)
		break
	default:
		http.NotFound(w, r)
	}
}
