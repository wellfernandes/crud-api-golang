package database

import (
	"crud-api-golang/configs"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Driver
)

// NewConnection open a connection database
func NewConnection() (*sql.DB, error) {

	_, _, _, _, strConn, err := configs.GetConfigInfo()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", strConn)
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
