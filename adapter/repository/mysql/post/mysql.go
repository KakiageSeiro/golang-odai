package post

import (
	"context"
	"golang-odai/domain"
	"golang-odai/external/mysql"
	"golang-odai/usecase/repository"

	"github.com/jinzhu/gorm"
)

type PostImpl struct {
	db       *mysql.DB
	repoUser repository.User
}

func New(db *mysql.DB, repoUser repository.User) repository.Post {
	return &PostImpl{
		db:       db,
		repoUser: repoUser,
	}
}

func (i *PostImpl) FindByID(ctx context.Context, id uint32) (*domain.Post, error) {
	var post domain.Post
	if err := i.db.Open().Where("id = ?", id).First(&post).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, repository.NotFoundRecord
		}
		return nil, err
	}

	return &post, nil
}

func (i *PostImpl) FindAll(ctx context.Context) ([]domain.Post, error) {
	posts := make([]domain.Post, 0)
	i.db.Open().Find(&posts)

	return posts, nil
}

func (i *PostImpl) Create(ctx context.Context, post *domain.Post) error {
	i.db.Open().Create(&post)

	return nil
}
