/*
작성자 : 박현준
작성일 : 2022.03.19.

수정자 : 박현준
수정일 : 2022.04.22.

파일 설명
DB 연결
*/

package persistence

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func getConnection() *sql.DB {
	if db == nil {
		initConnection()
	} else {
		err := db.Ping()
		if err != nil {
			initConnection()
		}
	}
	return db
}

func initConnection() *sql.DB {
	db_info := os.Getenv("BBADDABOTDB")
	var err error
	db, err = sql.Open("mysql", db_info)
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(3)
	db.SetConnMaxLifetime(time.Hour)
	if err != nil {
		fmt.Println("DB connect error")
		panic(err.Error())
	}
	return db
}
