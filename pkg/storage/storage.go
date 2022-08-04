package storage

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetMySQL(db *sql.DB, dsn string) error {
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}
