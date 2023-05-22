package database

import (
	"crud-api-golang/configs"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Driver
)

// Connect open a connection database
func NewConnection() (*sql.DB, error) {

	cfg := configs.DBConfig{}

	db, err := sql.Open("mysql", cfg.StrConn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
