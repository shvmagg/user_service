package db

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func DB() *sql.DB {
	return db
}

func Connect() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func Close() {
	db.Close()
}
