package post

import (
	post2 "golang-odai/adapter/http/render/post"
	"golang-odai/adapter/http/session"
	"golang-odai/usecase/interactor/post"
	"golang-odai/usecase/repository"

	"github.com/unrolled/render"
)

type Post struct {
	re       *post2.Render
	repoPost repository.Post
	sess     *session.Session
	usePost  post.Post
}

func New(sess *session.Session, re *render.Render, p repository.Post, usecase post.Post) *Post {
	return &Post{
		re:       post2.New(re),
		repoPost: p,
		sess:     sess,
		usePost:  usecase,
	}
}
