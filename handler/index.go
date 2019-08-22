package handler

import (
	"golang-odai/session"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/unrolled/render"
	"golang-odai/model"
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

	log.Printf(username)
	log.Printf(password)

	//ユーザーテーブルにユーザ名が存在する場合ログインできる
	postResult, err := model.IsLogin(r.Context(), username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if postResult != nil {
		log.Printf("Login Success!")

		s := &session.Data1{
			UserID: int(postResult.ID),
		}
		//セッションに保存
		err := session.SetData1(s, r ,w)
		if err != nil {
			panic(err)
		}

		//インデックス画面にリダイレクト
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		log.Printf("Login Failed.")
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

//post詳細画面
func PostDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	//投稿を取得
	post, err := model.FindByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//投稿に紐づくコメントを取得
	comment, err := model.CommentFindByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	re := render.New(render.Options{
		Charset: "UTF-8",
		Extensions: []string{".html"},
	})




	deta := struct {
		PostID uint32
		UserID uint32
		PostText string
		CommentText []model.Comment
	}{
		PostID: post.ID,
		UserID: post.UserID,
		PostText: post.Text,
		CommentText: comment,
	}

	re.HTML(w, http.StatusOK, "detail", deta)
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	re := render.New(render.Options{
		Charset: "UTF-8",
		Extensions: []string{".html"},
	})
	re.HTML(w, http.StatusOK, "form", nil)
}

// TODO: バリデーション追加する

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	ps, err := model.PasswordHash(password)
	if err != nil {
		panic(err)
	}

	p := model.User{
		Username: username,
		Password: ps,
	}

	if err := model.InsertUser(r.Context(), p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//セッションIDを生成してIDをDBに保持
	// sID, _ := uuid.NewV4()
	// c := &http.Cookie{
	// 	Name:  "session",
	// 	Value: sID.String(),
	// }
	// http.SetCookie(w, c)
	//TODO:ここにセッションテーブルにID入れる処理
	//dbSessions[c.Value] = un

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("CreateHandler")

	//ログインしているかを確認する
	data, err := session.GetData1(r)
	if err != nil {
		panic(err)
	}
	//ユーザーがテーブルに存在することを確認
	log.Printf(string(data.UserID))

	_, err = model.FindByUserId(r.Context(), data.UserID)
	if err != nil {
		panic(err)
	}

	log.Printf("投稿時ログイン確認OK")

	text := r.FormValue("text")

	p := model.Post{
		UserID: uint32(data.UserID),
		Text: text,
	}

	if err := model.Insert(r.Context(), p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//コメント投稿
func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("CreateCommentHandler")

	//ログインしているかを確認する
	data, err := session.GetData1(r)
	if err != nil {
		panic(err)
	}
	//ユーザーがテーブルに存在することを確認
	log.Printf(string(data.UserID))

	_, err = model.FindByUserId(r.Context(), data.UserID)
	if err != nil {
		panic(err)
	}

	log.Printf("コメント投稿時ログイン確認OK")

	text := r.FormValue("text")
	strPostID := r.FormValue("postID")
	postID, err := strconv.Atoi(strPostID)

	c := model.Comment{
		UserID: uint32(data.UserID),
		PostID: uint32(postID),
		Text: text,
	}

	if err := model.InsertComment(r.Context(), c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
