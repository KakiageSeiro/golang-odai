package authentication

import (
	"context"
	"golang-odai/domain"
	"golang-odai/external/firebase"
	"golang-odai/usecase/repository"
)

type (
	// SignInResponse /verifyPasswordのレスポンス
	SignInResponse struct {
		Kind         string `json:"kind"`
		IDToken      string `json:"idToken"`
		Email        string `json:"email"`
		DisplayName  string `json:"displayName"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    int    `json:"expiresIn,string"`
		LocalID      string `json:"localId"`
		Registered   bool   `json:"registered"`
	}

	// SignupNewUserResponse /signupNewUserのレスポンス
	SignupNewUserResponse struct {
		Kind         string `json:"kind"`
		IDToken      string `json:"idToken"`
		Email        string `json:"email"`
		RefreshToken string `json:"resreshToken"`
		ExpiresIn    int    `json:"expiresIn,string"`
		LocalID      string `json:"localId"`
	}

	// UserData signup/signin時にPOSTするdata
	UserData struct {
		Email             string `json:"email"`
		Password          string `json:"password"`
		ReturnSecureToken bool   `json:"returnSecureToken"`
	}
)

type AuthenticationImpl struct {
	f *firebase.Firebase
}

func New(f *firebase.Firebase) repository.Authentication {
	return &AuthenticationImpl{
		f: f,
	}
}

func (i *AuthenticationImpl) FindByEmailAndPassword(ctx context.Context, email, password string) (*domain.Authentication, error) {
	data := &UserData{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}

	var res SignInResponse
	if err := i.f.Post(ctx, "verifyPassword", data, &res); err != nil {
		return nil, err
	}

	return &domain.Authentication{
		ID: res.LocalID,
		Email: res.Email,
	}, nil
}

func (i *AuthenticationImpl) Create(ctx context.Context, email, password string)  (*domain.Authentication, error) {
	data := &UserData{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}

	var res SignupNewUserResponse
	if err := i.f.Post(ctx, "signupNewUser", data, &res); err != nil {
		return nil, err
	}

	return &domain.Authentication{
		ID: res.LocalID,
		Email: res.Email,
	}, nil
}
