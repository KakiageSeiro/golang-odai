package main

import (
	// "github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"golang-odai/handler"
	"html/template"
	"io"
)

func main() {

	e := echoInit()

	e.GET("/", handler.IndexHandler) //一覧画面
	e.GET("/signup", handler.SignupFormHandler) //ユーザー新規登録画面
	e.POST("/create_user", handler.CreateUserHandler) //ユーザー作成
	e.GET("/login", handler.LoginFormHandler) //ログイン画面
	e.POST("/login", handler.LoginHandler) //ログイン実行


	// //TODO:u.Use（ミドルウェア）を利用してログイン確認を共通処理とする
	//
	// e.GET("/posts/{id}", handler.PostDetailHandler)
	// e.GET("/form", handler.FormHandler)
	// e.POST("/create", handler.CreateHandler)
	// e.POST("/comment", handler.CreateCommentHandler)
	//

	e.Logger.Fatal(e.Start(":80"))
}

func echoInit() *echo.Echo {
	//Echoにテンプレートファイルを登録
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Renderer = t
	return e
}


//echoでのHTMLレンダリング用
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}