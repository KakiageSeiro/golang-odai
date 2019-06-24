package model

import (
	"context"
	"my-golang-odai/db"
	"strconv"
	"testing"
)

//ツイート
type post struct {
	Id   int
	Name string
	Text string
}

//ツイート投稿
func TestTweet(t *testing.T) {
	//テストデータ作成

	//ランダム文字列とか

	//呼び出し

	//検証
}

//一覧取得
func TestRetrieveLists(t *testing.T) {
	//テストデータ作成
	posts := make([]post, 0)
	for i := 0; i > 10; i++ {
		p := post{0, "ユーザー" + strconv.Itoa(i+1), "ツイート内容" + strconv.Itoa(i+1)}
		posts = append(posts, p)
	}
	insertPosts(nil, t, posts)

	//呼び出し
	lists, e := RetrieveLists(nil)
	if e != nil {
		t.Error("テスト対象関数からエラーが返却された")
	}

	//検証
	for _, v := range lists {
		pass := existInExpectedValues(posts, v)
		if !pass {
			t.Error("取得した値が期待値の中に含まれていません")
		}
	}
}
//
////一覧取得
//func TestRetrieveOneTweet(t *testing.T) {
//
//	//呼び出し
//	lists, e := RetrieveOneTweet(nil, 1)
//	if e != nil {
//		t.Error("テスト対象関数からエラーが返却された")
//	}
//
//	//検証
//	for _, v := range lists {
//		pass := existInExpectedValues(posts, v)
//		if !pass {
//			t.Error("取得した値が期待値の中に含まれていません")
//		}
//	}
//}
//
////期待値の中に同じデータが存在する場合true
//func existInExpectedValues(exp []post, act Post) bool {
//	for _, exp := range exp {
//		if act.Name == exp.Name && act.Text == exp.Text {
//			return true
//		}
//	}
//
//	return false
//}
//
////Postsテーブルにデータ追加
//func insertPosts(ctx context.Context, t *testing.T, posts []post) {
//	//DBコネクション取得
//
//	db, e := db.GetConnection("root@tcp(127.0.0.1:43306)/twitter")
//	if e != nil {
//		t.Error("DB接続で失敗")
//	}
//
//	//insert
//	for _, v := range posts {
//		if _, e := db.Open().QueryContext(
//			ctx,
//			"insert into posts(name, text) value(?, ?)",
//			v.Name,
//			v.Text,
//		); e != nil {
//			//失敗した場合
//			t.Error("テスト用データのinsertで失敗")
//		}
//	}
//}
