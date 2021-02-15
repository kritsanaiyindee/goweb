package main

import (
	"log"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./tpl/*.html",
	}

	rnd = renderer.New(opts)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	port := ":9000"
	log.Println("Listening on port ", port)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("static/assets"))))
	mux.Handle("/buttons/", http.StripPrefix("/buttons/", http.FileServer(http.Dir("static/buttons"))))
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	mux.Handle("/icons/", http.StripPrefix("/icons/", http.FileServer(http.Dir("static/icons"))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	mux.Handle("/vendors/", http.StripPrefix("/vendors/", http.FileServer(http.Dir("static/vendors"))))
	mux.Handle("/base/", http.StripPrefix("/base/", http.FileServer(http.Dir("static/base"))))
	http.ListenAndServe(port, mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "home", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "about", nil)
}

func cards(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "cards", nil)
}