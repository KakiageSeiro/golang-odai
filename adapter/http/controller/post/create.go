package post

import (
	"golang-odai/domain"
	"net/http"
)

func (hp *Post) Create(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")

	u, err := hp.sess.GetUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	p := &domain.Post{
		UserID: u.ID,
		Text:   text,
	}

	if err := hp.repoPost.Create(r.Context(), p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
