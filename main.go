package main

import (
	"fmt"
	"golang-api/controller"
	"golang-api/server"
	"net/http"

	"github.com/rs/cors"
)

var er = new(server.ErrorHandle)

type Mux struct{}

func main() {
	//Cors跨域設定
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:4200",
			"http://localhost:58505",
		},
	})

	fmt.Println("Cors:", "Ok")

	mux := &Mux{}
	//監聽Port設定
	err := http.ListenAndServe(":8778", c.Handler(mux))
	er.CheckErr(err)
}

//路由
func (m Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Port:", r.Host)
	fmt.Println("Path:", r.URL.Path)
	switch r.URL.Path {
	case "/CreateUser":
		controller.CreateUser(w, r)
	case "/UsersList":
		controller.UsersList(w, r)
	case "/DockerTest":
		controller.DockerTest(w, r)
	case "/DeleteUser":
		controller.DeleteUser(w, r)
	case "/ChannlTest":
		controller.ChannelTest(w, r)
	case "/SocketPage":
		controller.SockretPage(w, r)
	default:
		http.NotFound(w, r)
	}
}
