/*
작성자 : 박현준
작성일 : 2022.03.19.

수정자 : 박현준
수정일 : 2022.04.22.

파일 설명
History(No, UserNum, BeforeChannelId, AfterChannelId, Time, HistoryType) 테이블 매핑
*/

package persistence

import (
	ds "bbaddabot/datastruct"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// history 삽입 후 auto_increment 반환
func InsertHistory(h ds.History) (int, error) {
	db := getConnection()
	err := db.Ping()
	var id int64
	if err == nil {
		sql := `INSERT INTO  history (userNum, beforeChannelId, afterChannelId, time, historyType) 
		VALUES(?, ?, ?, now(), ?)`
		stmt, _ := db.Prepare(sql)
		res, _ := stmt.Exec(h.UserNum, h.BeforeChannelId, h.AfterChannelId, h.HistoryType)
		id, err = res.LastInsertId()
		stmt.Close()
	}
	return int(id), err
}

// 당일 공부 history 목록 조회
func SelectTodayHistoryByUserNum(userNum int) []ds.History {
	db := getConnection()
	err := db.Ping()
	var historys []ds.History
	if err == nil {

		sql := `SELECT no, username, beforechannelid, afterchannelid, time 
		FROM history
		WHERE userNum = ? AND DATE(time) = DATE(NOW())`
		stmt, _ := db.Prepare(sql)
		res, _ := stmt.Query(userNum)

		var history ds.History

		for res.Next() {
			res.Scan(&history.No, &history.UserNum, &history.BeforeChannelId, &history.AfterChannelId, &history.Time)
			historys = append(historys, history)
		}
		stmt.Close()
	}
	return historys
}

// 최근 채널에 머무른 시간 조회
func SelectMinuteSpentByUserNum(userNum int) int {
	db := getConnection()
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
	return studyTime
}
