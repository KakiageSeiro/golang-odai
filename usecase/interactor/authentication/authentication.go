package authentication

import (
	"context"
	"golang-odai/domain"
	"golang-odai/usecase/repository"
)

// Authentication defines the interface for authentication processing.
type Authentication interface {
	Register(ctx context.Context, name, email, password string) error
	Verify(ctx context.Context, email, password string) (*domain.User, error)
}

// authentication implements processing based on Interface.
type authentication struct {
	repoUser repository.User
	repoAuth repository.Authentication
}

// New returns new Authentication
func New(repoUser repository.User, repoAuth repository.Authentication) Authentication {
	return &authentication{
		repoUser: repoUser,
		repoAuth: repoAuth,
	}
}

func (a *authentication) Register(ctx context.Context, name, email, password string) error {
	auth, err := a.repoAuth.Create(ctx, email, password)
	if err != nil {
		return err
	}

	u := &domain.User{
		Name: name,
		AuthenticationID: auth.ID,
	}

	if _, err := a.repoUser.Create(ctx, u); err != nil {
		return err
	}

	return nil
}

func (a *authentication) Verify(ctx context.Context, email, password string) (*domain.User, error) {
	auth, err := a.repoAuth.FindByEmailAndPassword(ctx, email, password)
	if err != nil {
		return nil, err
	}

	u, err := a.repoUser.FindByAuthenticationID(ctx, auth.ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}
