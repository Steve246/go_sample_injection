package config

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Db *sqlx.DB
}

type dbConfig struct {
	dbHost string 
	dbPort string 
	dbName string 
	dbUser string 
	dbPassword string 
	dbDriver string 
}

func (c *Config) initDb() {
	var dbConfig = dbConfig{}
	dbConfig.dbHost = os.Getenv("DB_HOST")
	dbConfig.dbPort = os.Getenv("DB_PORT")
	dbConfig.dbUser = os.Getenv("DB_USER")
	dbConfig.dbPassword = os.Getenv("DB_PASSWORD")
	dbConfig.dbName = os.Getenv("DB_NAME")
	dbConfig.dbDriver = os.Getenv("DB_DRIVER")



	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", 
	dbConfig.dbDriver,
	dbConfig.dbUser,
	dbConfig.dbPassword,
	dbConfig.dbHost,
	dbConfig.dbPort,
	dbConfig.dbName,
)

db, err := sqlx.Connect(dbConfig.dbDriver, dsn)

if err != nil {
	panic(err)
}

c.Db = db 

}


func NewConfig() Config {
	cfg := Config{}
	cfg.initDb()
	return cfg
}