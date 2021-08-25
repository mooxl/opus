package main

import (
	"net/http"

	"github.com/gorilla/mux"

	c "opus/controller"
)

// NewRouter :
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static"))

	r.PathPrefix("/css/").Handler(fs)
	r.PathPrefix("/js/").Handler(fs)
	r.PathPrefix("/img/").Handler(fs)
	r.PathPrefix("/font/").Handler(fs)
	r.PathPrefix("/music/").Handler(fs)

	r.HandleFunc("/", c.HomeHandler)
	r.HandleFunc("/musician", c.MusicianHandler)

	return r
}

func main() {
	r := NewRouter()
	http.ListenAndServe(":443", r)
}
