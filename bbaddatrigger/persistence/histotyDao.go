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
	err := db.Ping()
	var id int64
	if err == nil {
		sql := `INSERT INTO  history (userNum, beforeChannelId, afterChannelId, time, historyType) 
		VALUES(?, ?, ?, now(), ?)`
		// msg := fmt.Sprintf("insert : %#v", h)
		// fmt.Println(msg)
		stmt, _ := db.Prepare(sql)
		res, _ := stmt.Exec(h.UserNum, h.BeforeChannelId, h.AfterChannelId, h.HistoryType)
		id, err = res.LastInsertId()
		stmt.Close()
	}
	defer db.Close()
	return int(id), err
}

// 오늘 공부 기록 조회해 목록 반환
func SelectTodayHistoryByUserNum(userNum int) []ds.History {
	db := dbConn()
	err := db.Ping()
	var historys []ds.History
	if err == nil {

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
		stmt.Close()
	}
	defer db.Close()
	return historys
}

// 혼자쓰면 이렇게 가능한데
// 여러명 동시에 사용하는거 고려하면 이거만으로는 안된다..
// 이번 해당 history 를 기준으로 채널에 있었던 분 확인
//
func SelectMinuteSpentChannel(historyNo int) int {
	db := dbConn()
	err := db.Ping()
	var studyTime int
	if err == nil {
		sql := `SELECT TIMESTAMPDIFF(MINUTE,
			(SELECT time FROM history
			WHERE no = ? -1),
			(SELECT time FROM history
			WHERE no = ?
			))`
		stmt, _ := db.Prepare(sql)
		stmt.QueryRow(historyNo, historyNo).Scan(&studyTime)
		stmt.Close()
	}
	defer db.Close()
	return studyTime
}

// UserNum 을 받아서 최신의 값 2개만 가져오고 그 값으로 비교해보자.
func SelectMinuteSpentChannel2(userNum int) int {
	db := dbConn()
	err := db.Ping()
	var studyTime int
	if err == nil {
		sql := `
		SELECT TIMESTAMPDIFF(MINUTE,
			(select time
			from (SELECT time, rank() over (order by time desc) as 'rank'
					FROM history
					where userNum = ? and date(time) = date(now())) as recent
			where recent.rank =2),
			(select time
			from (SELECT time, rank() over (order by time desc) as 'rank'
					FROM history
					where userNum = ? and date(time) = date(now())) as recent
			where recent.rank =1)) as time
	`

		stmt, _ := db.Prepare(sql)
		stmt.QueryRow(userNum, userNum).Scan(&studyTime)
		stmt.Close()
	}
	defer db.Close()
	return studyTime
}
