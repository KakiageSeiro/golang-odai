package post

import (
	"golang-odai/adapter/http/render/post"
	"golang-odai/domain"
	"net/http"
)

type Data struct {
	Posts []domain.Post
}

func (p *Post) Index(w http.ResponseWriter, r *http.Request) {
	posts, err := p.repoPost.FindAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := p.re.Index(w, post.IndexData{
		Posts: posts,
	}); err != nil {
		panic(err)
	}
}
