package handler

import (
	"github.com/go-chi/chi"
	"github.com/unrolled/render"
	"my-golang-odai/model"
	"net/http"
	"strconv"
)

//ツイート一覧
type Data struct{
	List []model.Post
}

//ツイート一覧画面
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//ツイートをすべて取得
	list, err := model.RetrieveLists(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//標準ライブラリ利用
	//// テンプレート読み込み
	//t := templates.Must(templates.ParseFiles("templates/Index.html"))
	//// テンプレートを描画
	//if err := t.Execute(w, list); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}

	//render利用
	data := Data{
		list,
	}
	renderer(w, "index", data)
}

//単一のツイート表示画面
func OneTweetHandler(w http.ResponseWriter, r *http.Request) {
	//クエリパラメータ取得
	id := chi.URLParam(r, "id")
	if id == "" {
		//TODO:新しくエラーを作って返したいね
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

	//標準ライブラリ利用
	// テンプレート読み込み
	//t := template.Must(template.ParseFiles("templates/tweet.html"))
	//// テンプレートを描画
	//if err := t.Execute(w, *post); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}

	//render利用
	renderer(w, "tweet", post)
}

//投稿（ツイート）用フォーム画面
func FormHandler(w http.ResponseWriter, r *http.Request) {

	//標準ライブラリ利用
	//// テンプレート読み込み
	//t := template.Must(template.ParseFiles("templates/form.html"))
	//// テンプレートを描画
	//if err := t.Execute(w, nil); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}

	//render利用
	renderer(w, "form", nil)
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

	//TODO:この方法だとHTTPヘッダに余計な書き込みをしているメッセージが出る（HTMLレンダリング自体はできてる）
	//render利用
	//re := render.New(render.Options{
	//	Extensions: []string{".html"},
	//	Charset:    "UTF-8",
	//})
	//re.HTML(w, http.StatusSeeOther, "index", nil)
}

//render利用でHTMLレンダリングする
func renderer(w http.ResponseWriter, templateName string, data interface{}) {
	re := render.New(render.Options{
		Extensions: []string{".html"},
		Charset:    "UTF-8",
	})
	re.HTML(w, http.StatusOK, templateName, data)
}
