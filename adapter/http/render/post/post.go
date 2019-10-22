package post

import (
	"golang-odai/domain"
	"net/http"
	"path"

	"github.com/unrolled/render"
)

type Render struct {
	re *render.Render
}

type IndexData struct {
	Posts []domain.Post
}

type DetailData struct {
	Post *domain.Post
}

type FormData struct {
	User *domain.User
}

const TEMPLATE_DIR = "post"

func New(re *render.Render) *Render {
	return &Render{
		re: re,
	}
}

func (r *Render) Index(w http.ResponseWriter, data IndexData) error {
	return r.re.HTML(w, http.StatusOK, path.Join(TEMPLATE_DIR, "index"), data)
}

func (r *Render) Detail(w http.ResponseWriter, data DetailData) error {
	return r.re.HTML(w, http.StatusOK, path.Join(TEMPLATE_DIR, "detail"), data)
}

func (r *Render) Form(w http.ResponseWriter, data FormData) error {
	return r.re.HTML(w, http.StatusOK, path.Join(TEMPLATE_DIR, "form"), data)
}
