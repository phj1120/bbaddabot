package persistence

import (
	ds "bbaddabot/datastruct"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// no, username, before_channel, after_channel, time

// history 삽입 후 auto_increment 반환
func InsertHistory(h ds.History) (int, error) {
	db := dbConn()
	sql := `INSERT INTO  history (userNum, beforeChannelId, afterChannelId, time, historyType) 
	VALUES(?, ?, ?, now(), ?)`
	stmt, _ := db.Prepare(sql)
	res, _ := stmt.Exec(h.UserNum, h.BeforeChannelId, h.AfterChannelId, h.HistoryType)
	id, err := res.LastInsertId()

	return int(id), err
}

// 오늘 공부 기록 조회해 목록 반환
func SelectTodayHistoryByUserNum(userNum int) []ds.History {
	db := dbConn()
	sql := `SELECT *
			FROM history
			WHERE userNum = ? AND DATE(time) = DATE(NOW())`
	stmt, _ := db.Prepare(sql)
	res, _ := stmt.Query(userNum)
	var historys []ds.History
	var history ds.History

	for res.Next() {
		res.Scan(&history.No, &history.UserNum, &history.BeforeChannelId, &history.AfterChannelId, &history.Time, &history.Time)
		historys = append(historys, history)
	}

	return historys
}

// 혼자쓰면 이렇게 가능한데
// 여러명 동시에 사용하는거 고려하면 이거만으로는 안된다..
// 이번 해당 history 를 기준으로 채널에 있었던 분 확인
//
func SelectMinuteSpentChannel(historyNo int) int {
	db := dbConn()
	sql := `SELECT TIMESTAMPDIFF(MINUTE,
		(SELECT time FROM history
		WHERE no = ? -1),
		(SELECT time FROM history
		WHERE no = ?
		))`
	stmt, _ := db.Prepare(sql)
	var studyTime int
	stmt.QueryRow(historyNo, historyNo).Scan(&studyTime)
	return studyTime
}
