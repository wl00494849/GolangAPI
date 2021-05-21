package main

import (
	"fmt"
	"golang-api/controller"
	"golang-api/server"
	"net/http"

	"golang.org/x/net/websocket"
)

var er = new(server.ErrorHandle)

type Mux struct{}

func main() {
	//Cors跨域設定
	// c := cors.New(cors.Options{
	// 	AllowedOrigins: []string{
	// 		"http://localhost:4200",
	// 		"http://localhost:58505",
	// 	},
	// })

	fmt.Println("Cors:", "Ok")
	http.HandleFunc("/SocketPage", controller.SockretPage)
	http.Handle("/Socket", websocket.Handler(server.Echo))
	http.ListenAndServe(":8778", nil)

	// mux := &Mux{}
	// //監聽Port設定
	// err := http.ListenAndServe(":8778", c.Handler(mux))
	// er.CheckErr(err)

	// http.Handle("/", websocket.Server{})
	// http.ListenAndServe(":8877", nil)

}

//路由
func (m Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/CreateUser":
		controller.CreateUser(w, r)
		break
	case "/UsersList":
		controller.UsersList(w, r)
		break
	case "/DockerTest":
		controller.DockerTest(w, r)
		break
	case "/DeleteUser":
		controller.DeleteUser(w, r)
		break
	case "/ChannlTest":
		controller.ChannelTest(w, r)
		break
	default:
		http.NotFound(w, r)
	}
}
