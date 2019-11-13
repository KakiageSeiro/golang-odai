package main

import (
	// "github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"golang-odai/handler"
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {

	// r := chi.NewRouter()

	e := echo.New()

	// メモ：handlerのパッケージ分離したほうが見通し良くなる
	// Goそのものよりレイヤーアーキテクチャ
	// e.GET()
	e.GET("/", handler.IndexHandler)

	// e.GET("/signup", handler.SignupFormHandler)
	//
	//
	// // r.Post("/signup", handler.SignupHandler)
	//
	// e.GET("/login", handler.LoginFormHandler)
	// e.POST("/login", handler.LoginHandler)
	//
	// //TODO:u.Use（ミドルウェア）を利用してログイン確認を共通処理とする
	//
	// e.GET("/posts/{id}", handler.PostDetailHandler)
	// e.GET("/form", handler.FormHandler)
	// e.POST("/create", handler.CreateHandler)
	// e.POST("/comment", handler.CreateCommentHandler)
	//
	// e.POST("/create_user", handler.CreateUserHandler)

	http.ListenAndServe(":80", e)
}