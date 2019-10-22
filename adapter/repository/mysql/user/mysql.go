package user

import (
	"context"
	"golang-odai/domain"
	"golang-odai/external/mysql"
	"golang-odai/usecase/repository"

	"github.com/jinzhu/gorm"
)

type UserImpl struct {
	db *mysql.DB
}

func New(db *mysql.DB) repository.User {
	return &UserImpl{db: db}
}

func (i *UserImpl) FindByIDs(_ context.Context, id ...uint32) ([]domain.User, error) {
	us := []domain.User{}
	if err := i.db.Open().Where("id IN (?)", id).Find(&us).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, repository.NotFoundRecord
		}
		return nil, err
	}

	return us, nil
}

func (i *UserImpl) FindByAuthenticationID(_ context.Context, authID string) (*domain.User, error) {
	u := domain.User{}
	if err := i.db.Open().Where("authentication_id = ?", authID).First(&u).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, repository.NotFoundRecord
		}
		return nil, err
	}

	return &u, nil
}

func (i *UserImpl) FindByID(_ context.Context, id uint32) (*domain.User, error) {
	u := domain.User{}
	if err := i.db.Open().Where("id = ?", id).First(&u).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, repository.NotFoundRecord
		}
		return nil, err
	}

	return &u, nil
}

func (i *UserImpl) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	i.db.Open().Create(&u)

	return u, nil
}
