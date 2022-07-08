package config

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Db *sqlx.DB
}

type ApiConfig struct {
	ApiHost string 
	ApiPort string 
}

type dbConfig struct {
	dbHost string 
	dbPort string 
	dbName string 
	dbUser string 
	dbPassword string 
	dbDriver string 
}

// DB_DRIVER=postgres DB_HOST=localhost DB_PORT=5432 DB_USER=postgres DB_PASSWORD= DB_NAME=postgres API_HOST=localhost API_PORT=8888 go run .


func (c *Config) initDb() {
	var dbConfig = dbConfig{}
	dbConfig.dbHost = os.Getenv("DB_HOST")
	dbConfig.dbPort = os.Getenv("DB_PORT")
	dbConfig.dbUser = os.Getenv("DB_USER")
	dbConfig.dbPassword = os.Getenv("DB_PASSWORD")
	dbConfig.dbName = os.Getenv("DB_NAME") //db_sql_injection
	dbConfig.dbDriver = os.Getenv("DB_DRIVER") //postgres



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

func (c *Config) DbConn() *sqlx.DB {
	return c.Db
}




func NewConfig() Config {
	cfg := Config{}
	cfg.initDb()
	return cfg
}