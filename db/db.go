package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//接続先
const dsn = "root@tcp(db)/twitter"

// DB database interface
type DB struct {
	conn *sql.DB
}

//DB接続し、コネクションを取得
func GetConnection() (*DB, error) {
	//DB接続
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db := DB{conn: conn}

	//コネクションの最大ライフタイム（タイムアウト時間）
	conn.SetConnMaxLifetime(10 * time.Second)
	//コネクションの最大コネクション数
	conn.SetMaxOpenConns(10)
	//コネクションの最大待機数
	conn.SetMaxIdleConns(10)

	//接続確認
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return &db, nil
}

// Open returns the database connection.
func (d *DB) Open() *sql.DB {
	return d.conn
}
