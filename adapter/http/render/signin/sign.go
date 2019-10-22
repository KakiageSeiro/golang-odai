package signin

import (
	"net/http"
	"path"

	"github.com/unrolled/render"
)

type Render struct {
	re *render.Render
}

const TEMPLATE_DIR = "signin"

func New(re *render.Render) *Render {
	return &Render{
		re: re,
	}
}

func (r *Render) Form(w http.ResponseWriter) error {
	return r.re.HTML(w, http.StatusOK, path.Join(TEMPLATE_DIR, "form"), nil)
}
