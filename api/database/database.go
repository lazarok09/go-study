package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dataSourceName := "lazarok09:lazarok09@/devbook"
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		return nil, err
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, err
}
