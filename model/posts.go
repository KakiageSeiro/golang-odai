package model

import (
	"context"
	"my-golang-odai/db"
)

//ツイートのデータ型
type Post struct {
	ID   int
	Name string
	Text string
}

//すべての投稿を取得
func RetrieveLists(ctx context.Context) ([]Post, error) {

	//DB接続
	db, e := db.GetConnection()
	if e != nil {
		return nil, e
	}

	//すべての投稿を取得
	rows, e := db.Open().QueryContext(
		ctx,
		"select name, text from posts",
	)

	//リストにして返却
	list := make([]Post, 0)
	for rows.Next() {
		//１レコード取り出してリストに追加
		var p Post
		if err := rows.Scan(&p.Name, &p.Text); err != nil {
			return nil, err
		}
		list = append(list, p)
	}

	return list, nil
}
