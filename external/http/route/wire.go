//+build wireinject

package route

import (
	"github.com/google/wire"
	"golang-odai/adapter/http/controller/index"
	"golang-odai/adapter/http/controller/signin"
	"golang-odai/adapter/http/controller/signup"
	"golang-odai/adapter/http/render"
	"golang-odai/adapter/http/session"
	"golang-odai/external/firebase"
	"golang-odai/external/mysql"
	"golang-odai/adapter/http/controller/post"
	"net/http"
)

func BuildIndexController(db *mysql.DB, fb *firebase.Firebase, r *render.Config, s *session.Config) *index.Index {
	wire.Build(SuperSet, DomainSet, ControllerSet)
	return nil
}

func BuildPostController(db *mysql.DB, fb *firebase.Firebase, r *render.Config, s *session.Config) *post.Post {
	wire.Build(SuperSet, DomainSet, UsecaseSet, ControllerSet)
	return nil
}

func BuildSignupController(db *mysql.DB, fb *firebase.Firebase, r *render.Config, s *session.Config) *signup.Signup {
	wire.Build(SuperSet, DomainSet, UsecaseSet, ControllerSet)
	return nil
}

func BuildSigninController(db *mysql.DB, fb *firebase.Firebase, r *render.Config, s *session.Config) *signin.Sign {
	wire.Build(SuperSet, DomainSet, UsecaseSet, ControllerSet)
	return nil
}

func BuildAuthenticationMiddleware(fb *firebase.Firebase, s *session.Config) func(http.Handler) http.Handler {
	wire.Build(SuperSet, MiddlewareSet)
	return nil
}