package index

import (
	index2 "golang-odai/adapter/http/render/index"
	"golang-odai/usecase/interactor/timeline"
	"golang-odai/usecase/repository"
	"net/http"

	"github.com/unrolled/render"
)

type Index struct {
	re          *index2.Render
	repoPost    repository.Post
	useTimeline timeline.TimeLine
}

func New(re *render.Render, p repository.Post, usecase timeline.TimeLine) *Index {
	return &Index{
		re:          index2.New(re),
		repoPost:    p,
		useTimeline: usecase,
	}
}

func (i *Index) Index(w http.ResponseWriter, r *http.Request) {
	posts, err := i.useTimeline.GetLatestList(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := i.re.Index(w, index2.Data{
		Posts: posts,
	}); err != nil {
		panic(err)
	}
}
