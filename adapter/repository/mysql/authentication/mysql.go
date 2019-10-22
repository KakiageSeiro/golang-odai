package authentication

import (
	"context"
	"golang-odai/domain"
	"golang-odai/external/mysql"
	"golang-odai/usecase/repository"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationImpl struct {
	db *mysql.DB
}

func New(db *mysql.DB) repository.Authentication {
	return &AuthenticationImpl{db: db}
}

func (i *AuthenticationImpl) FindByEmailAndPassword(_ context.Context, email, password string) (*domain.Authentication, error) {
	a := &domain.Authentication{}
	if err := i.db.Open().Where("email = ?", email).First(&a).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, repository.NotFoundRecord
		}
		return nil, err
	}

	// パスワードのハッシュ検査
	// DBから取得したハッシュ化済みパスワードと、フォームから取得した生パスワードが同等のものか検査
	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password)); err != nil {
		return nil, err
	}

	return a, nil
}

func (i *AuthenticationImpl) Create(ctx context.Context, email, password string)  (*domain.Authentication, error) {

	a := &domain.Authentication{
		Email: email,
	}

	// パスワードを不可逆なハッシュ化
	hashPW, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	a.Password = string(hashPW)

	i.db.Open().Create(&a)

	return a, nil
}
