package main

import (
	"net/http"
	"golang-odai/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", handler.IndexHandler)

	r.Get("/loginform", handler.LoginFormHandler)
	r.Post("/login", handler.LoginHandler)

	r.Get("/posts/{id}", handler.PostDetailHandler)
	r.Get("/form", handler.FormHandler)
	r.Post("/create", handler.CreateHandler)
	http.ListenAndServe(":80", r)
}
