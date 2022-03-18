package persistence

import (
	ds "bbaddabot/datastruct"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// userNum, studyTime, date

// 공부 시간 생성(당일 처음 공부 시작한 경우 )
func InsertNewStudyTotal(userNum int, studyTime int) int {
	db := dbConn()
	sql := `INSERT INTO  studyTotal (userNum, studyTime, date) 
	VALUES(?, ?, DATE(now()));`

	stmt, _ := db.Prepare(sql)
	res, _ := stmt.Exec(userNum, studyTime)
	id, _ := res.LastInsertId()

	return int(id)
}

// 오늘 공부 시간 조회
func SelectStudyTotalTodayByUserNum(userNum int) (int, error) {
	db := dbConn()
	sql := `SELECT studyTime
			FROM studyTotal
			WHERE date=DATE(NOW()) AND userNum=?`
	stmt, _ := db.Prepare(sql)

	// 행 하나 : QueryRow
	var studyTime int
	err := stmt.QueryRow(userNum).Scan(&studyTime)

	return studyTime, err
}

// 공부 시간 업데이트 후 영향 받은 행 개수 반환
func UpdateStudyTimeByUserNumAndStudyTime(userNum int, studyTime int) int {
	db := dbConn()
	sql := `UPDATE studyTotal 
			SET studytime = studytime + ?
			WHERE date=DATE(NOW()) AND userNum=?`
	stmt, _ := db.Prepare(sql)
	res, _ := stmt.Exec(studyTime, userNum)

	cnt, _ := res.RowsAffected()

	return int(cnt)
}

func DeletestudyTotal(h ds.StudyTotal) (*int64, error) {
	return nil, nil

}
