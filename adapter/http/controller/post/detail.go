package post

import (
	post2 "golang-odai/adapter/http/render/post"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (p *Post) Detail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	post, err := p.usePost.Get(r.Context(), uint32(uid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := p.re.Detail(w, post2.DetailData{Post: post}); err != nil {
		panic(err)
	}
}
