package model

import (
	"context"
	"github.com/pkg/errors"
	"github.com/jinzhu/gorm"
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
	// メモ：gormのメソッドチェーン使うといい
	if err := db.Open().Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		if (gorm.IsRecordNotFoundError(err)) {
			return false, NotFoundRecord
		}
		return false, err
	}

	//レコードが習得できなかった場合はログイン不可
	// メモ：stringにnilは入らない。ポインタならOK
	// if user.Username == "" {
	// 	return false, nil
	// }
	// メモ：こっちがいい
	if user == nil {
		return false, nil
	}

	//レコードをしゅとくできたらログイン可能とみなす
	return true, err
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