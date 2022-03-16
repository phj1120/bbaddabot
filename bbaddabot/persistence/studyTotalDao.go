package persistence

import (
	ds "bbaddabot/datastruct"
	_ "database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// No, Username, Study_time, Study_time

func InsertstudyTotal(h ds.StudyTotal) (*int64, error) {
	db := dbConn()
	sql := `insert into study_total(username, study_time, date) 
			values (?, ?, DATE(now()))`

	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(h.Username, h.Study_time, h.Date)
	if err != nil {
		return nil, err
	}

	id, e := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &id, e
}

func SelectTodaystudyTotal(userName string) ([]ds.studyTotal, error) {
	db := dbConn()
	sql := `SELECT * 
			FROM studyTotal 
			WHERE username = ? AND DATE(time) = DATE(NOW())`
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Query(userName)
	if err != nil {
		return nil, err
	}

	var studyTotals []ds.studyTotal
	var studyTotal ds.studyTotal

	for res.Next() {
		err := res.Scan(&studyTotal.No, &studyTotal.Username, &studyTotal.Before_channel, &studyTotal.After_channel, &studyTotal.Time)
		if err != nil {
			return nil, err
		}

		fmt.Println(studyTotal)
		studyTotals = append(studyTotals, studyTotal)
	}
	return studyTotals, nil
}

func UpdatestudyTotal(h ds.studyTotal) (*int64, error) {
	return nil, nil
}

func DeletestudyTotal(h ds.studyTotal) (*int64, error) {
	return nil, nil

}
