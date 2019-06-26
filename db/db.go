package db

import (
	//"database/sql" //標準ライブラリ利用
	//_ "github.com/go-sql-driver/mysql"

	_ "github.com/go-sql-driver/mysql"
	//Gorm利用
	"github.com/jinzhu/gorm"
	"time"
)

//接続先
//Docker-composeでのmy-golang-odai_db_1のように、dbという文字列が存在していれば以下のようにtcp(db)でOK
const DSN = "root@tcp(db)/twitter"

// DB database interface
type DB struct {
	//conn *sql.DB
	conn *gorm.DB
}

//DB接続し、コネクションを取得
func GetConnection() (*DB, error) {

	//DB接続
	//if dsn == ""{ //テスト環境で接続先を変えようとした残滓
	//	dsn = DSN
	//}

	//標準ライブラリ利用
	//conn, err := sql.Open("mysql", DSN)
	//if err != nil {
	//	return nil, err
	//}
	//db := DB{conn: conn}

	//Gorm利用
	conn, err := gorm.Open("mysql", DSN)
	if err != nil {
		panic("DBに接続できませんでした。。。■" + err.Error())
	}
	//defer conn.Close()
	db := DB{conn: conn}

	//コネクションの最大ライフタイム（タイムアウト時間）
	conn.DB().SetConnMaxLifetime(10 * time.Second)
	//コネクションの最大コネクション数
	conn.DB().SetMaxOpenConns(10)
	//コネクションの最大待機数
	conn.DB().SetMaxIdleConns(10)

	//接続確認
	if err := conn.DB().Ping(); err != nil {
		return nil, err
	}

	return &db, nil
}

// Open returns the database connection.
func (d *DB) Open() *gorm.DB {
	return d.conn
}
