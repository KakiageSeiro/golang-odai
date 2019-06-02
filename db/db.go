package db

import (
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

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

	conn.SetConnMaxLifetime(10 * time.Second)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return &db, nil
}

// Open returns the database connection.
func (d *DB) Open() *sql.DB {
	return d.conn
}
