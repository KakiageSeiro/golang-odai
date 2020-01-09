package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/labstack/echo/v4"
	"github.com/unrolled/render"
	"golang-odai/model"
	"golang-odai/session"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	apiURI   = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/%s?key=%s"
	tokenURI = "https://securetoken.googleapis.com/v1/token?key=%s"
)

const APIKEY = "AIzaSyBxDwSVBvI-j49sj9lOOkVEsuF00LsEx-Q"

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

type Data struct {
	Posts []model.Post
}

//サインアップフォーム
func SignupFormHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "signup", nil)
}

//ログインフォーム
func LoginFormHandler(c echo.Context) error{
	return c.Render(http.StatusOK, "login", nil)
}

//ログイン実行
func LoginHandler(c echo.Context) error {
	//パラメータ取得
	username := c.Param("username")
	password := c.Param("password")

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
	//s := &session.Data1{
	//	SessionID: localID,
	//}

	////セッションに保存
	//err3 := session.SetData1(s, r, w)
	//if err3 != nil {
	//	panic(err3)
	//}

	//インデックス画面にリダイレクト
	//http.Redirect(w, r, "/", http.StatusSeeOther)
	return c.Redirect(http.StatusMovedPermanently, "/")
}

//firebaseでログイン情報を検証
func LoginVerify(username, password string) (string, error) {
	println(username)
	println(password)

	data := UserData{
		Email:             username,
		Password:          password,
		ReturnSecureToken: true,
	}

	var res SignInResponse
	if err := post(context.Background(), "verifyPassword", data, APIKEY, &res); err != nil {
		panic(err)
	}

	//ログインできた場合はセッションに格納
	log.Printf("Login Success!")

	log.Printf("%+v¥n", res)
	return res.LocalID, nil
}






















func IndexHandler(c echo.Context) error {
	posts, err := model.Select()
	if err != nil {
		echo.NewHTTPError(http.StatusNotFound)
	}

	log.Print(posts)


	return c.Render(http.StatusOK, "index", posts)
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
		Charset:    "UTF-8",
		Extensions: []string{".html"},
	})

	deta := struct {
		PostID      uint32
		UserID      string
		PostText    string
		CommentText []model.Comment
	}{
		PostID:      post.ID,
		UserID:      post.UserID,
		PostText:    post.Text,
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
		Charset:    "UTF-8",
		Extensions: []string{".html"},
	})
	re.HTML(w, http.StatusOK, "form", nil)
}

//ユーザー新規作成
func CreateUserHandler(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	println(username)
	println(password)

	data := UserData{
		Email:             username,
		Password:          password,
		ReturnSecureToken: false,
	}

	//Firebaseにユーザー作成
	var res SignupNewUserResponse
	if err := post(c.Request().Context(), "signupNewUser", data, APIKEY, &res); err != nil {
		panic(err)
	}

	//セッションIDとユーザーネームをDBに保存
	//ログイン用
	user := model.User{
		SessionID: res.LocalID,
		Username:  username,
	}
	model.InsertUser(user)

	//インデックス画面にリダイレクト
	//http.Redirect(w, r, "/login", http.StatusSeeOther)
	return c.Redirect(http.StatusMovedPermanently, "/login")
}

//Firebaseにユーザー作成
func post(ctx context.Context, service string, data interface{}, apikey string, resp interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	//Firebaseへのリクエスト作成
	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(apiURI, service, apikey),
		strings.NewReader(string(b)),
	)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(request.WithContext(ctx))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Read the response body
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, res.Body); err != nil {
		return err
	}

	println("ここかな")
	if res.StatusCode == http.StatusBadRequest {
		panic(res)
	}
	println("ここかも")

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
		Text:   text,
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
		Text:   text,
	}

	if err := model.InsertComment(r.Context(), c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
