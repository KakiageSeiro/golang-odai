package handler

import (
	"fmt"
	"net/http"

	"golang-odai/model"
	"github.com/go-chi/chi"
	"github.com/unrolled/render"

)

type Data struct{
	Posts []model.Post
}

func IndexRender(w http.ResponseWriter,posts []model.Post) {
	re := render.New(render.Options{
		Charset: "UTF-8",
		Extensions: []string{".html"},
	})
	data := Data{
		posts,
	}
	re.HTML(w, http.StatusOK, "index", data)
}

//ログインフォーム
func SignupFormHandler(w http.ResponseWriter, r *http.Request) {
	re := render.New(render.Options{
		Charset: "UTF-8",
		Extensions: []string{".html"},
	})
	re.HTML(w, http.StatusOK, "signup", nil)
}

//ログイン実行
// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	//パラメータ取得
// 	username := r.FormValue("username")
// 	password := r.FormValue("password")
//
// 	//ユーザーテーブルにユーザ名が存在する場合ログインできる
// 	postResult, err := model.IsLogin(r.Context(), username, password)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
//
// 	if postResult {
// 		fmt.Fprintln(w, "Login Success!")
//
// 		//インデックス画面にリダイレクト
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 	} else {
// 		fmt.Fprintln(w, "Login Failed.")
// 	}
// }

//ログインフォーム
func LoginFormHandler(w http.ResponseWriter, r *http.Request) {
	re := render.New(render.Options{
		Charset: "UTF-8",
		Extensions: []string{".html"},
	})
	re.HTML(w, http.StatusOK, "login", nil)
}

//ログイン実行
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	//パラメータ取得
	username := r.FormValue("username")
	password := r.FormValue("password")

	//ユーザーテーブルにユーザ名が存在する場合ログインできる
	postResult, err := model.IsLogin(r.Context(), username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if postResult {
		fmt.Fprintln(w, "Login Success!")

		//インデックス画面にリダイレクト
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		fmt.Fprintln(w, "Login Failed.")
	}
}


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := model.Select(r.Context())
	if err != nil {
		/*
		if err == model.Notfound {
			not found
		}
		*/

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	IndexRender(w, posts)
}

func PostDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	post, err := model.FindByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	re := render.New(render.Options{
		Charset: "UTF-8",
		Extensions: []string{".html"},
	})
	re.HTML(w, http.StatusOK, "detail", post)
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	re := render.New(render.Options{
		Charset: "UTF-8",
		Extensions: []string{".html"},
	})
	re.HTML(w, http.StatusOK, "form", nil)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	p := model.User{
		Username: username,
		Password: password,
	}

	if err := model.InsertUser(r.Context(), p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// func CreateHandler(w http.ResponseWriter, r *http.Request) {
// 	name := r.FormValue("name")
// 	text := r.FormValue("text")
//
// 	p := model.Post{
// 		Name: name,
// 		Text: text,
// 	}
//
// 	if err := model.Insert(r.Context(), p); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
//
// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }