package handler

import (
	"github.com/go-chi/chi"
	"html/template"
	"my-golang-odai/model"
	"net/http"
	"strconv"
)

//ツイート一覧画面
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//ツイートをすべて取得
	lists, err := model.RetrieveLists(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// テンプレート読み込み
	t := template.Must(template.ParseFiles("template/Index.html"))
	// テンプレートを描画
	if err := t.Execute(w, lists); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//単一のツイート表示画面
func OneTweetHandler(w http.ResponseWriter, r *http.Request) {
	//クエリパラメータ取得
	id := chi.URLParam(r, "id")
	if id == "" {
		//TODO:新しくエラーを作って返したいけどよくわからん
		//http.Error(w, errors.New(""), http.StatusInternalServerError)
	}

	//ツイートを１件取得
	idInt, e := strconv.Atoi(id)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}

	post, err := model.RetrieveOneTweet(r.Context(), idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// テンプレート読み込み
	t := template.Must(template.ParseFiles("template/tweet.html"))
	// テンプレートを描画
	if err := t.Execute(w, *post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//投稿（ツイート）用フォーム画面
func FormHandler(w http.ResponseWriter, r *http.Request) {
	// テンプレート読み込み
	t := template.Must(template.ParseFiles("template/form.html"))
	// テンプレートを描画
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//投稿（ツイート）する
func TweetHandler(w http.ResponseWriter, r *http.Request) {
	//POSTのみ許可（そもそもハンドラ作成時点でメソッドまで指定してルーティングしたい）
	if r.Method != http.MethodPost {
		//405を返す
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("POST以外ｗｗｗありえないｗｗｗ"))
		return
	}

	//パラメータ取得
	r.ParseForm()      //Bodyデータを扱う場合には、事前にパースを行う
	form := r.PostForm //Formデータを取得

	//投稿
	model.Tweet(r.Context(), form.Get("name"), form.Get("text"))

	//リダイレクト(index行き)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
