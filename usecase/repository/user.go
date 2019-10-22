package repository

import (
	"context"
	"golang-odai/domain"

	"github.com/pkg/errors"
)

var NotFoundRecord = errors.New("Notfound")

type User interface {
	FindByAuthenticationID(ctx context.Context, authID string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	FindByIDs(ctx context.Context, id ...uint32) ([]domain.User, error)
	FindByID(ctx context.Context, id uint32) (*domain.User, error)
}
