package configs

import (
	"crud-api-golang/internal/constants"
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

	cfg.Port, err = strconv.Atoi(os.Getenv(constants.API_PORT))
	if err != nil {
		cfg.Port = constants.DEFAULT_API_PORT
	}

	cfg.Host = os.Getenv(constants.DATABASE_HOST)
	cfg.User = os.Getenv(constants.DATABASE_USER)
	cfg.Password = os.Getenv(constants.DATABASE_PASSWORD)
	cfg.Database = os.Getenv(constants.DATABASE_NAME)

	cfg.StrConn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Database)

	return nil
}

// GetConfigInfo returns config infos
func GetConfigInfo() (int, string, string, string, string, error) {
	if cfg.Port == 0 || cfg.Host == "" ||
		cfg.Database == "" || cfg.User == "" || cfg.StrConn == "" {
		return 0, "", "", "", "", errors.New(constants.ERROR_ENVIROMENT_VARIABLES)
	}

	return cfg.Port, cfg.Host, cfg.Database, cfg.User, cfg.StrConn, nil
}
