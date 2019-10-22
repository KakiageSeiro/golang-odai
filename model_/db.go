package model

import (
	"time"
	"github.com/jinzhu/gorm"
)

const dsn = "root@tcp(db)/twitter"

// DB database interface
type DB struct {
	conn *gorm.DB
}

// NewDB is DB constructor.
func New() (*DB, error) {

	conn, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db := DB{conn: conn}

	conn.DB().SetConnMaxLifetime(10 * time.Second)
	conn.DB().SetMaxOpenConns(10)
	conn.DB().SetMaxIdleConns(10)

	if err := conn.DB().Ping(); err != nil {
		return nil, err
	}
	return &db, nil
}

// Open returns the database connection.
func (d *DB) Open() *gorm.DB {
	return d.conn
}
