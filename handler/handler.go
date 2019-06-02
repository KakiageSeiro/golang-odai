package handler

import (
	"fmt"
	"my-golang-odai/model"
	"net/http"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "★★★IndexHandler１")

	lists, err := model.RetrieveLists(r.Context())
	fmt.Fprintln(w, "★★★取得したツイート数■" + strconv.Itoa(len(lists)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintln(w, "★★★IndexHandler２")

	//テスト用描画
	for _, e := range lists {
		//デバッグ用ログ
		fmt.Fprintln(w, "★★★ツイートユーザー名■" + e.Name)
		fmt.Fprintln(w, "★★★ツイート内容■" + e.Text)
		fmt.Fprintln(w, "■■■■■■■■■■■■■■■■■■■■■■■")
	}




}
