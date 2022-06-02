package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	driver   = os.Getenv("DB_DRIVER")
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	dbName   = os.Getenv("DB_NAME")
	db       = &sql.DB{}
)

func init() {
	var err error
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if db.Ping() != nil {
		panic("could not connect to database")
	}

}

func GetDB() *sql.DB {
	return db
}
