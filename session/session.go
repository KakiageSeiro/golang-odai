package session

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"io"
	"strings"
)

//initという名称にするとパッケージが初めてインポートされたタイミングで実行される
func init(){

	// 構造体を登録
	gob.Register(&Data1{})
	// セッション初期処理
	sessionInit()

}

// セッション名
var session_name string = "gsid"
// Cookie型のstore情報
var strore *sessions.CookieStore
// セッションオブジェクト
var session *sessions.Session

// 構造体
type Data1 struct {
	Count    int
	Msg      string
}


// セッション用の初期処理
func sessionInit(){

	// 乱数生成
	b := make([]byte, 48)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		panic(err)
	}
	str := strings.TrimRight(base32.StdEncoding.EncodeToString(b), "=")

	// 新しいstoreとセッションを準備
	store = sessions.NewCookieStore([]byte(str))
	session = sessions.NewSession(store, session_name)

	// セッションの有効範囲を指定
	store.Options = &sessions.Options{
		Domain:     "localhost",
		Path:       "/",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   true,
	}

	// log
	fmt.Println("key     data --")
	fmt.Println(str)
	fmt.Println("")
	fmt.Println("store   data --")
	fmt.Println(store)
	fmt.Println("")
	fmt.Println("session data --")
	fmt.Println(session)
	fmt.Println("")

}

func GetData1()(*Data1, error){

	data1, ok := session.Values["data1"].(*Data1)
	if !ok {
		return nil, errors.New("セッション取得失敗")
	}

	if data1 == nil {
		return nil, errors.New("セッションの中身がなかった")
	}
	return data1, nil
}

func SetData1(data1 *Data1)(){

}