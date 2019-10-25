package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/unrolled/render"
	"golang-odai/model"
	"golang-odai/session"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"context"
)

var (
	apiURI = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/%s?key=%s"
	tokenURI = "https://securetoken.googleapis.com/v1/token?key=%s"
)

type (
	// SignInResponse /verifyPasswordのレスポンス
	SignInResponse struct {
		Kind         string `json:"kind"`
		IDToken      string `json:"idToken"`
		Email        string `json:"email"`
		DisplayName  string `json:"displayName"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    int    `json:"expiresIn,string"`
		LocalID      string `json:"localId"`
		Registered   bool   `json:"registered"`
	}

	// SignupNewUserResponse /signupNewUserのレスポンス
	SignupNewUserResponse struct {
		Kind         string `json:"kind"`
		IDToken      string `json:"idToken"`
		Email        string `json:"email"`
		RefreshToken string `json:"resreshToken"`
		ExpiresIn    int    `json:"expiresIn,string"`
		LocalID      string `json:"localId"`
	}

	// UserData signup/signin時にPOSTするdata
	UserData struct {
		Email             string `json:"email"`
		Password          string `json:"password"`
		ReturnSecureToken bool   `json:"returnSecureToken"`
	}
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

//サインアップフォーム
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

	//firebaseでユーザーとパスの検証をする
	localID, err2 := LoginVerify(username, password)
	if err2 != nil {
		log.Printf("Login Failed.")
	}

	log.Printf("Login Success!")

	log.Println("■■■ｓ")
	log.Println(localID)
	s := &session.Data1{
		SessionID: localID,
	}

	//セッションに保存
	err3 := session.SetData1(s, r ,w)
	if err3 != nil {
		panic(err3)
	}

	//インデックス画面にリダイレクト
	http.Redirect(w, r, "/", http.StatusSeeOther)

}


//firebaseでログイン情報を検証
func LoginVerify(username, password string) (string, error){
	println(username)
	println(password)

	data := UserData{
		Email:             username,
		Password:          password,
		ReturnSecureToken: true,
	}

	var res SignInResponse
	if err := post(context.Background(), "verifyPassword", data, "AIzaSyAS_a8LX-EhpVqD_7rQALPMjViGc_NPpI8", &res); err != nil {
		panic(err)
	}

	//ログインできた場合はセッションに格納
	log.Printf("Login Success!")

	log.Printf("%+v¥n", res)
	return res.LocalID, nil
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
		UserID string
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

	//ログインしているかを確認する
	data, err := session.GetData1(r)
	log.Println(data.SessionID)
	if err != nil {
		panic(err)
	}


	re := render.New(render.Options{
		Charset: "UTF-8",
		Extensions: []string{".html"},
	})
	re.HTML(w, http.StatusOK, "form", nil)
}


//ユーザー新規作成
func CreateUserHandler(w http.ResponseWriter, r *http.Request)  {
	username := r.FormValue("username")
	password := r.FormValue("password")
	println(username)
	println(password)

	data := UserData{
		Email:             username,
		Password:          password,
		ReturnSecureToken: false,
	}

	//Firebaseにユーザー作成
	var res SignupNewUserResponse
	//if err := post(context.Background(), "signupNewUser", data, "AIzaSyAS_a8LX-EhpVqD_7rQALPMjViGc_NPpI8", &res); err != nil {
	//	panic(err)
	//}

	if err := post(context.Background(), "signupNewUser", data, "AIzaSyAS_a8LX-EhpVqD_7rQALPMjViGc_NPpI8", &res); err != nil {
		panic(err)
	}


	//セッションIDとユーザーネームをDBに保存
	//ログイン用
	user := model.User {
		SessionID: res.LocalID,
		Username: username,
	}
	model.InsertUser(r.Context(), user)





	//TODO:結果を確認しエラーだった場合はエラーページ行き


	//インデックス画面にリダイレクト
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}

//Firebaseにユーザー作成
func post(ctx context.Context, service string, data interface{}, apikey string, resp interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	//Firebaseへのリクエスト作成
	r, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(apiURI, service, apikey),
		strings.NewReader(string(b)),
	)
	if err != nil {
		return err
	}

	r.Header.Set("Content-Type", "application/json")


	client := &http.Client{}
	res, err := client.Do(r.WithContext(ctx))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Read the response body
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, res.Body); err != nil {
		return err
	}

	if res.StatusCode == http.StatusBadRequest {
		panic(res)
	}

	return json.Unmarshal(buf.Bytes(), &resp)
}

//投稿
func CreateHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("■■■■■■CreateHandler")

	//ログインしているかを確認する
	data, err := session.GetData1(r)
	log.Println(data.SessionID)
	if err != nil {
		panic(err)
	}


	user, err := model.FindBySessionId(r.Context(), data.SessionID)
	if err != nil {
		panic(err)
	}

	log.Printf("投稿時ログイン確認OK")

	text := r.FormValue("text")

	p := model.Post{
		UserID: user.Username,
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

	user, err := model.FindBySessionId(r.Context(), data.SessionID)
	if err != nil {
		panic(err)
	}

	log.Printf("コメント投稿時ログイン確認OK")

	text := r.FormValue("text")
	strPostID := r.FormValue("postID")
	postID, err := strconv.Atoi(strPostID)

	c := model.Comment{
		UserID: user.Username,
		PostID: uint32(postID),
		Text: text,
	}

	if err := model.InsertComment(r.Context(), c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}