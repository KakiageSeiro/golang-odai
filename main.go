package main

import (
	"github.com/go-chi/chi"
	"my-golang-odai/handler"
	"net/http"
)

func main() {
	//標準ライブラリVer
	//http.HandleFunc("/", handler.IndexHandler)
	//http.HandleFunc("/form", handler.FormHandler)
	//http.HandleFunc("/tweet", handler.TweetHandler)

	//Chi利用Ver
	r := chi.NewRouter()
	r.Get("/", handler.IndexHandler) //一覧
	r.Get("/form", handler.FormHandler) //ツイート用フォーム
	r.Post("/tweet", handler.TweetHandler) //ツイートする
	//単一ツイート取得
	r.Route("/posts", func(r chi.Router) {
		r.Get("/{id}", handler.OneTweetHandler)

		//r.Use(ArticleCtx)
		//r.Get("/", getArticle)                                          // GET /articles/123
		//r.Put("/", updateArticle)                                       // PUT /articles/123
		//r.Delete("/", deleteArticle)                                    // DELETE /articles/123
	})

	http.ListenAndServe(":80", r)
}
