package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySqlStorage(cfq mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfq.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
