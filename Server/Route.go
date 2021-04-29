package Server

import (
	"log"
	"net/http"
	"text/template"
)

type Route struct{}

func (r Route) RedirectRoute(w http.ResponseWriter, url string) {
	t, _ := template.ParseFiles(url)
	log.Println(t.Execute(w, nil))
}
