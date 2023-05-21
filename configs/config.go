package configs

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type DBConfig struct {
	StrConn  string
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

var cfg DBConfig

// Load loads the environment variables
func Load() error {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg.Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		cfg.Port = 9000 // default port
	}

	cfg.Host = os.Getenv("HOST")
	cfg.User = os.Getenv("DATABASE_USER")
	cfg.Password = os.Getenv("DATABASE_PASSWORD")
	cfg.Database = os.Getenv("DATABASE_NOME")

	cfg.StrConn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Database)

	return nil
}

func GetConfigInfo() (int, string, string, string, error) {
	if cfg.Port == 0 || cfg.Host == "" ||
		cfg.Database == "" || cfg.StrConn == "" {
		return 0, "", "", "", errors.New("error in environment variables")
	}

	return cfg.Port, cfg.Host, cfg.Database, cfg.StrConn, nil
}
