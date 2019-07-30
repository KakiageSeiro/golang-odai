package main

import (
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"golang-odai/handler"
	"net/http"
)

func main() {

	r := chi.NewRouter()

	// メモ：handlerのパッケージ分離したほうが見通し良くなる
	// Goそのものよりレイヤーアーキテクチャ

	r.Get("/", handler.IndexHandler)

	r.Get("/signup", handler.SignupFormHandler)
	// r.Post("/signup", handler.SignupHandler)

	r.Get("/login", handler.LoginFormHandler)
	r.Post("/login", handler.LoginHandler)

	//TODO:u.Use（ミドルウェア）を利用してログイン確認を共通処理とする

	r.Get("/posts/{id}", handler.PostDetailHandler)
	r.Get("/form", handler.FormHandler)
	r.Post("/create", handler.CreateHandler)

	r.Post("/create_user", handler.CreateUserHandler)

	http.ListenAndServe(":80", r)
}
