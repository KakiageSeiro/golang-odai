package repository

import (
	"context"
	"golang-odai/domain"
)

type Authentication interface {
	FindByEmailAndPassword(ctx context.Context, email, password string) (*domain.Authentication, error)
	Create(ctx context.Context, email, password string) (*domain.Authentication, error)
}
