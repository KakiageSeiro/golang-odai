package signup

import (
	signup2 "golang-odai/adapter/http/render/signup"
	"golang-odai/usecase/interactor/authentication"
	"golang-odai/usecase/repository"

	"github.com/unrolled/render"
)

type Signup struct {
	re       *signup2.Render
	repoUser repository.User
	useAuth authentication.Authentication
}

func New(re *render.Render, p repository.User, useAuth authentication.Authentication) *Signup {
	return &Signup{
		re:       signup2.New(re),
		repoUser: p,
		useAuth: useAuth,
	}
}
