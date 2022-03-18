package persistence

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	db_info := os.Getenv("BBADDABOTDB")
	db, err := sql.Open("mysql", db_info)
	if err != nil {
		panic(err.Error())
	}
	return db
}
