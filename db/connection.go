package db

import (
	"crud-api-golang/configs"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //postgresql driver
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgress", sc)
	if err != nil {
		panic(err) //it is not recommended to use panic in production
	}
	err = conn.Ping()

	return conn, err
}
