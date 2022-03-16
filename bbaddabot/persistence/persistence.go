package persistence

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func PersistenceTest() {
	fmt.Println("PersistenceTest")
}

func SqlTest() {
	// sql.DB 객체 생성
	db, err := sql.Open("mysql", "bbadda:mybbadda@tcp(127.0.0.1:3306)/bbaddabot")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 하나의 Row를 갖는 SQL 쿼리
	var bbadda int
	// err = db.QueryRow("SELECT bbadda FROM channel_total WHERE username='park'").Scan(&bbadda)
	err = db.QueryRow("insert into channel_total values ('kk', 3)").Scan(&bbadda)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bbadda)
}
