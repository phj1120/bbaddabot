package persistence

// import (
// 	ds "bbaddabot/datastruct"
// 	_ "database/sql"
// 	"fmt"

// 	_ "github.com/go-sql-driver/mysql"
// )

// // no, username, before_channel, after_channel, time
// func InsertHistory(h ds.History) (*int64, error) {
// 	db := dbConn()
// 	sql := `insert into history(username, before_channel, after_channel, time)
// 			values (?, ?, ?, Now())`
// 	stmt, _ := db.Prepare(sql)
// 	res, _ := stmt.Exec(h.UserNum, h.BeforeChannelId, h.AfterChannelId, h.HistoryType)

// 	id, err := res.LastInsertId()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &id, e
// }

// func SelectTodayHistory(userName string) ([]ds.History, error) {
// 	db := dbConn()
// 	sql := `SELECT *
// 			FROM history
// 			WHERE username = ? AND DATE(time) = DATE(NOW())`
// 	stmt, err := db.Prepare(sql)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res, err := stmt.Query(userName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var historys []ds.History
// 	var history ds.History

// 	for res.Next() {
// 		err := res.Scan(&history.No, &history.Username, &history.Before_channel, &history.After_channel, &history.Time)
// 		if err != nil {
// 			return nil, err
// 		}

// 		fmt.Println(history)
// 		historys = append(historys, history)
// 	}
// 	return historys, nil
// }

// func UpdateHistory(h ds.History) (*int64, error) {
// 	return nil, nil
// }

// func DeleteHistory(h ds.History) (*int64, error) {
// 	return nil, nil

// }
