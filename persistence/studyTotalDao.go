/*
작성자 : 박현준
작성일 : 2022.03.19.

수정자 : 박현준
수정일 : 2022.04.22.

파일 설명
studyTotal(no, userNum, studyTime, date, WeekSuccessCnt, TodaySuccess) 테이블 매핑
*/

package persistence

import (
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// 공부 시간 생성(당일 처음 공부 시작한 경우 )
func InsertNewStudyTotalByUserNumAndStudyTime(userNum int, studyTime int) int {
	db := getConnection()
	err := db.Ping()
	var id int64
	if err == nil {
		sql := `INSERT INTO  studyTotal (userNum, studyTime, date) 
		VALUES(?, ?, DATE(now()));`

		stmt, _ := db.Prepare(sql)
		res, _ := stmt.Exec(userNum, studyTime)
		id, _ = res.LastInsertId()
		stmt.Close()
	}
	return int(id)
}

// userNum 으로 오늘 공부 시간 조회
func SelectStudyTotalTodayByUserNum(userNum int) (int, error) {
	db := getConnection()
	err := db.Ping()
	var studyTime int
	if err == nil {
		sql := `SELECT studyTime
		FROM studyTotal
		WHERE date=DATE(NOW()) AND userNum=?`

		stmt, _ := db.Prepare(sql)
		err = stmt.QueryRow(userNum).Scan(&studyTime)
		stmt.Close()
	}
	return studyTime, err
}

// userId 와 guildId로 오늘 공부 시간 조회
func SelectStudyTotalTodayByUserIdAndGuildId(userId string, guildId string) (int, error) {
	db := getConnection()
	err := db.Ping()
	var studyTime int
	if err == nil {
		sql := `select studytime from bbaddabot.studyTotal
		WHERE date =DATE(NOW()) and 
		usernum = (select usernum from user 
			where userid=? and guildid = ?)`

		stmt, _ := db.Prepare(sql)
		err = stmt.QueryRow(userId, guildId).Scan(&studyTime)
		stmt.Close()
	}
	return studyTime, err

}

// 이번 주 공부 시간 조회
func SelectStudyTotalWeekByUserIdAndGuildId(userId string, guildId string) (int, error) {
	db := getConnection()
	err := db.Ping()
	var studyTime int
	if err == nil {
		sql := `SELECT sum(studyTime)
		FROM
			studyTotal
		WHERE
			date 
			BETWEEN
				(SELECT ADDDATE(CURDATE(), - WEEKDAY(CURDATE()) + 0 ))
			AND
				(SELECT ADDDATE(CURDATE(), - WEEKDAY(CURDATE()) + 6 ))
			AND
				usernum = (select usernum from user 
			where userid=? and guildid = ?)`

		stmt, _ := db.Prepare(sql)
		err = stmt.QueryRow(userId, guildId).Scan(&studyTime)
		stmt.Close()
	}
	return studyTime, err
}

// 이번 달 공부 시간 조회
func SelectStudyTotalMonthByUserIdAndGuildId(userId string, guildId string) (int, error) {
	db := getConnection()
	err := db.Ping()
	var studyTime int
	if err == nil {
		sql := `SELECT sum(studyTime)
		FROM
			studyTotal
		WHERE
			date 
			BETWEEN
				(SELECT DATE_FORMAT(NOW(), '%x-%m-01'))
			AND
				(SELECT LAST_DAY(NOW()))
			AND
				usernum = (select usernum from user 
			where userid=? and guildid = ?)`

		stmt, _ := db.Prepare(sql)
		err = stmt.QueryRow(userId, guildId).Scan(&studyTime)
		stmt.Close()
	}
	return studyTime, err
}

// 공부 시간 업데이트 후 영향 받은 행 개수 반환
func UpdateStudyTimeByUserNumAndStudyTime(userNum int, studyTime int) int {
	db := getConnection()
	err := db.Ping()
	var cnt int64
	if err == nil {
		sql := `UPDATE studyTotal 
		SET studytime = studytime + ?
		WHERE date=DATE(NOW()) AND userNum=?`
		stmt, _ := db.Prepare(sql)
		res, _ := stmt.Exec(studyTime, userNum)

		cnt, _ = res.RowsAffected()
		stmt.Close()
	}
	return int(cnt)
}
