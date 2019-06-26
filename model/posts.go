package model

import (
	"context"
	"my-golang-odai/db"

	//Gorm利用
	_ "github.com/jinzhu/gorm"
)

//ツイートのデータ型
type Post struct {
	ID   int
	Name string
	Text string
}

//すべての投稿を取得
func RetrieveLists(ctx context.Context) ([]Post, error) {
	//DBコネクション取得
	db, e := db.GetConnection()
	if e != nil {
		return nil, e
	}

	//すべての投稿を取得

	//標準ライブラリ利用
	//rows, e := db.Open().QueryContext(
	//	ctx,
	//	"select name, text from posts",
	//)
	//
	////リストにして返却
	//list := make([]Post, 0)
	//for rows.Next() {
	//	//１レコード取り出してリストに追加
	//	var p Post
	//	if err := rows.Scan(&p.Name, &p.Text); err != nil {
	//		return nil, err
	//	}
	//	list = append(list, p)
	//}

	//Gorm利用
	list := make([]Post, 0)
	db.Open().Find(&list)

	return list, nil
}

//投稿を１件取得
func RetrieveOneTweet(ctx context.Context, id int) (*Post, error) {
	//DBコネクション取得
	db, e := db.GetConnection()
	if e != nil {
		return nil, e
	}

	//投稿を取得
	//rows, e := db.Open().QueryContext(
	//	ctx,
	//	"select name, text from posts where id = ?",
	//	id,
	//	//strconv.Itoa(id),
	//)
	//if e != nil {
	//	return nil, e
	//}
	//
	////リストにして返却
	//list := make([]Post, 0)
	//var p Post
	//for rows.Next() {
	//	//１レコード取り出してリストに追加
	//	if err := rows.Scan(&p.Name, &p.Text); err != nil {
	//		return nil, err
	//	}
	//	list = append(list, p)
	//}
	//log.Printf("■ツイート一見取得処理で取得できたレコード数:" + strconv.Itoa(len(list)))

	//TODO:１レコードであることをチェックして、エラーの場合↓
	//NotFoundエラーを作って返す。呼び出し側でエラー種別をcaseしてハンドリングする

	//Gorm利用
	post := &Post{}
	db.Open().Where("id = ?", id).First(&post) //ポインタのポインタを渡すことに注意

	return post, nil
}

//ツイートする
func Tweet(ctx context.Context, name string, text string) error {
	//DBコネクション取得
	db, e := db.GetConnection()
	if e != nil {
		return e
	}

	//投稿内容を保存
	//if _, e := db.Open().QueryContext(
	//	ctx,
	//	"insert into posts(name, text) value(?, ?)",
	//	name,
	//	text,
	//); e != nil {
	//	//失敗した場合
	//	return e
	//}

	post := &Post{Name:name, Text:text}
	if result := db.Open().NewRecord(&post); result == true{
		db.Open().Create(&post)
	}

	return nil
}
