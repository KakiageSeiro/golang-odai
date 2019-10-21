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
<<<<<<< HEAD
	UserID uint32
	Text string
}

//コメント
type Comment struct {
	ID uint32
	UserID uint32
	PostID uint32
=======
	Name string
>>>>>>> 95a6fbd55e2ebbb8edba0eda21840f5764a6fe3a
	Text string
}

//ログイン用
type User struct {
<<<<<<< HEAD
	ID uint32
=======
>>>>>>> 95a6fbd55e2ebbb8edba0eda21840f5764a6fe3a
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

<<<<<<< HEAD
//ログインできる場合はユーザーIDを返す
func IsLogin(_ context.Context, username string, password string) (*User, error) {
	db, err := New()
	if err != nil {
		return nil, err
=======
//ログインできる場合はtrue
func IsLogin(_ context.Context, username string, password string) (bool, error) {
	db, err := New()
	if err != nil {
		return false, err
>>>>>>> 95a6fbd55e2ebbb8edba0eda21840f5764a6fe3a
	}

	// 初期値 nil
	user := &User{}
	if err := db.Open().Where("username = ?", username).First(&user).Error; err != nil {
		if (gorm.IsRecordNotFoundError(err)) {
<<<<<<< HEAD
			return nil, NotFoundRecord
		}
		return nil, err
=======
			return false, NotFoundRecord
		}
		return false, err
>>>>>>> 95a6fbd55e2ebbb8edba0eda21840f5764a6fe3a
	}

	log.Printf(user.Password)

	// メモ：ifのなかの変数定義はスコープがifのなかだけになるから、ほかと同じ変数名が使える
	if err := passwordVerify(user.Password, password); err != nil {
<<<<<<< HEAD
		return nil, err
=======
		return false, err
>>>>>>> 95a6fbd55e2ebbb8edba0eda21840f5764a6fe3a
	}

	println("認証しました")

	//レコードが習得できなかった場合はログイン不可
<<<<<<< HEAD
	if user == nil {
		return nil, nil
	}

	//レコードをしゅとくできたらログイン可能とみなす
	return user, err
=======
	return user != nil, nil
	// ↑ リファクタリング後
	// if user == nil {
	// 	return false, nil
	// }
	//
	// //レコードをしゅとくできたらログイン可能とみなす
	// return true, err
>>>>>>> 95a6fbd55e2ebbb8edba0eda21840f5764a6fe3a
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
<<<<<<< HEAD
}

func InsertComment(ctx context.Context, comment Comment) error {
	db, err := New()
	if err != nil {
		return err
	}
	db.Open().Create(&comment)

	return nil
}

func CommentFindByID(_ context.Context, postId string) ([]Comment, error) {
	db, err := New()
	if err != nil {
		return nil, err
	}

	var comment []Comment
	if err := db.Open().Where("post_id = ?", postId).Find(&comment).Error; err != nil {
		if (gorm.IsRecordNotFoundError(err)) {
			return nil, NotFoundRecord
		}
		return nil, err
	}

	return comment, nil
}

func FindByUserId(_ context.Context, id int) (*User, error) {
	db, err := New()
	if err != nil {
		return nil, err
	}

	user := &User{}
	if err := db.Open().Where("id = ?", id).First(&user).Error; err != nil {
		if (gorm.IsRecordNotFoundError(err)) {
			return nil, NotFoundRecord
		}
		return nil, err
	}

	return user, nil
=======
>>>>>>> 95a6fbd55e2ebbb8edba0eda21840f5764a6fe3a
}