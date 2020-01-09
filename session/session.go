package session

import (
	"encoding/gob"
	"errors"
	"github.com/gorilla/sessions"
	"net/http"
)

//initという名称にするとパッケージが初めてインポートされたタイミングで実行される
func init() {

	// 構造体を登録
	gob.Register(&Data1{})
	// セッション初期処理
	sessionInit()

}

// セッション名
var session_name string = "gsid"

// Cookie型のstore情報
var store *sessions.CookieStore

// 構造体
type Data1 struct {
	SessionID string
}

// セッション用の初期処理
func sessionInit() {

	//// 乱数生成
	//b := make([]byte, 48)
	//_, err := io.ReadFull(rand.Reader, b)
	//if err != nil {
	//	panic(err)
	//}
	//str := strings.TrimRight(base32.StdEncoding.EncodeToString(b), "=")

	// 新しいstoreとセッションを準備
	store = sessions.NewCookieStore([]byte("session"))
	//session = sessions.NewSession(store, session_name)

	// セッションの有効範囲を指定
	store.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   1000,
		Secure:   false,
		HttpOnly: true,
	}

	//// log
	//fmt.Println("key     data --")
	//fmt.Println(str)
	//fmt.Println("")
	//fmt.Println("store   data --")
	//fmt.Println(store)
	//fmt.Println("")
	//fmt.Println("session data --")
	//fmt.Println(session)
	//fmt.Println("")

}

func GetData1(r *http.Request) (*Data1, error) {
	session, _ := store.Get(r, session_name)

	data1, ok := session.Values["data1"].(*Data1)
	if !ok {
		return nil, errors.New("セッション取得失敗")
	}

	if data1 == nil {
		return nil, errors.New("セッションの中身がなかった")
	}
	return data1, nil
}

func SetData1(data1 *Data1, r *http.Request, w http.ResponseWriter) error {
	session, _ := store.Get(r, session_name)

	session.Values["data1"] = data1

	// 保存
	return sessions.Save(r, w)
}
