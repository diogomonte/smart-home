package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect(driverName, dataSourceName string) *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
