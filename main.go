package main

import (
	// "github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"golang-odai/handler"
	"html/template"
	"io"
	"net/http"
	"github.com/labstack/echo/v4"
)





//echoでのHTMLレンダリング用
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}

func main() {

	e := echo.New()

	//HTMLファイル指定
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/index.html")),
	}

	e.Renderer = t
	e.GET("/", Hello)















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

	//http.ListenAndServe(":80", e)
}