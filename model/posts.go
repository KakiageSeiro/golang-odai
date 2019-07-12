package model

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Post struct {
	ID uint32
	Name string
	Text string
}

//ログイン用
type User struct {
	Username string
	Password string
}

var NotFoundRecord = errors.New("Notfound")

func FindByID(_ context.Context, id string) (*Post, error) {
	db, err := New()
	if err != nil {
		return nil, err
	}

	post := &Post{}
	if err := db.Open().Where("id = ?", id).First(&post).Error; err != nil {
		if (gorm.IsRecordNotFoundError(err)) {
			return nil, NotFoundRecord
		}
		return nil, err
	}

	return post, nil
}

//ログインできる場合はtrue
func IsLogin(_ context.Context, username string, password string) (bool, error) {
	db, err := New()
	if err != nil {
		return false, err
	}

	// 初期値 nil
	user := &User{}
	if err := db.Open().Where("username = ?", username).First(&user).Error; err != nil {
		if (gorm.IsRecordNotFoundError(err)) {
			return false, NotFoundRecord
		}
		return false, err
	}

	log.Printf(user.Password)

	// メモ：ifのなかの変数定義はスコープがifのなかだけになるから、ほかと同じ変数名が使える
	if err := passwordVerify(user.Password, password); err != nil {
		return false, err
	}

	println("認証しました")

	//レコードが習得できなかった場合はログイン不可
	return user != nil, nil
	// ↑ リファクタリング後
	// if user == nil {
	// 	return false, nil
	// }
	//
	// //レコードをしゅとくできたらログイン可能とみなす
	// return true, err
}

// パスワードハッシュを作る
func PasswordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// パスワードがハッシュにマッチするか
func passwordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

func Select(_ context.Context) ([]Post, error) {
	db, err := New()
	if err != nil {
		return nil, err
	}

	posts := make([]Post, 0)
	db.Open().Find(&posts)

	return posts, nil
}

func InsertUser(ctx context.Context, user User) error {
	db, err := New()
	if err != nil {
		return err
	}
	db.Open().Create(&user)

	return nil
}

func Insert(ctx context.Context, post Post) error {
	db, err := New()
	if err != nil {
		return err
	}
	db.Open().Create(&post)

	return nil
}