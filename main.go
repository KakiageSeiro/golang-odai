package main

import (
	"github.com/go-chi/chi"
	"my-golang-odai/handler"
	"net/http"
)

func main() {
	//標準ライブラリVer
	//http.HandleFunc("/", handler.IndexHandler)
	//http.HandleFunc("/form", handler.FormHandler)
	//http.HandleFunc("/tweet", handler.TweetHandler)

	//Chi利用Ver
	r := chi.NewRouter()
	r.Get("/", handler.IndexHandler)
	r.Get("/form", handler.FormHandler)
	r.Post("/tweet", handler.TweetHandler)

	http.ListenAndServe(":80", r)
}
