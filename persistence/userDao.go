package persistence

import (
	ds "bbaddabot/datastruct"

	_ "github.com/go-sql-driver/mysql"
)

// userNum, userId, guildId, userName, userType, wantTime, wantCnt

// User 추가 (기본 공부 목표 주 5회 3시간)
func UserSave(u ds.User) (*int64, error) {
	db := getConnection()
	err := db.Ping()
	var id int64
	if err == nil {
		sql := `INSERT INTO  user (userId, guildId, userName, userType, wantTime, wantCnt) 
		VALUES(?, ?, ?, 'user', '180', '5')`

		stmt, _ := db.Prepare(sql)
		res, _ := stmt.Exec(u.UserId, u.GuildId, u.UserName)
		id, err = res.LastInsertId()
		stmt.Close()
	}
	return &id, err
}

// User 조회
func UserFindByUserNum(userNum int) (ds.User, error) {
	db := getConnection()
	err := db.Ping()
	var user ds.User
	if err == nil {
		sql := `SELECT userNum, userId, guildId, userName, userType, wantTime, wantCnt 
		FROM user WHERE userNum=?`
		stmt, _ := db.Prepare(sql)

		err = stmt.QueryRow(userNum).Scan(&user.UserNum, &user.UserId, &user.GuildId, &user.UserName, &user.UserType, &user.WantTime, &user.WantCnt)
		stmt.Close()
	}
	return user, err
}

// UserNum 조회
func SelectUserNumByUserIdAndGuildId(userId string, guildId string) (int, error) {
	db := getConnection()
	err := db.Ping()
	var userNum int64
	if err == nil {
		sql := `SELECT userNum FROM user WHERE userId=? AND guildId=?`
		stmt, _ := db.Prepare(sql)
		err = stmt.QueryRow(userId, guildId).Scan(&userNum)
		stmt.Close()
	}
	return int(userNum), err
}

// 해당 길드의 User 리스트 조회
func UserListFindByGuildId(guildId string) ([]ds.User, error) {
	db := getConnection()
	err := db.Ping()
	var user ds.User
	var users []ds.User
	if err == nil {
		sql := `SELECT userNum, userId, guildId, userName, userType, wantTime, wantCnt 
		from user where guildId = ?`
		stmt, _ := db.Prepare(sql)
		res, _ := stmt.Query(guildId)

		for res.Next() {
			res.Scan(&user.UserNum, &user.UserId, &user.GuildId, &user.UserName, &user.UserType, &user.WantTime, &user.WantCnt)
			users = append(users, user)
		}
		stmt.Close()
	}
	return users, err
}

func SelectGuildIdList() ([]string, error) {
	db := getConnection()
	err := db.Ping()
	var guildId string
	var guildIds []string
	if err == nil {
		sql := "select distinct guildid from user"
		stmt, _ := db.Prepare(sql)
		res, _ := stmt.Query()

		for res.Next() {
			res.Scan(&guildId)
			guildIds = append(guildIds, guildId)
		}
		stmt.Close()
	}
	return guildIds, err
}

// 유저 타입
func UpdateUserType(u ds.User) (*int64, error) {
	db := getConnection()
	err := db.Ping()
	var no int64
	if err == nil {
		sql := `UPDATE user SET type = ? 
		WHERE userNum = ?`
		stmt, _ := db.Prepare(sql)
		res, _ := stmt.Exec(u.UserType, u.UserId)
		// auto_increment 반환
		no, err = res.LastInsertId()
		stmt.Close()
	}
	return &no, err
}

// 사용자 삭제
func DeleteUser(userId string) (*int64, error) {
	db := getConnection()
	err := db.Ping()
	var no int64
	if err == nil {
		sql := `DELETE FROM user WHERE userId = ?`

		stmt, _ := db.Prepare(sql)
		res, _ := stmt.Exec(userId)
		// 반영된 row 수 반환
		no, err = res.RowsAffected()
		stmt.Close()
	}
	return &no, err
}

// SELECT 예시
// func SelectUserNumByUserIdAndGuildId(userId string, guildId string) (int, error) {

// 	db := getConnection()
// 	sql := `SELECT userNum
// 	FROM user WHERE userId=? AND guildId=?`
// 	stmt, _ := db.Prepare(sql)

// 	// 행 하나 : QueryRow
// 	var userNum int
// 	err := stmt.QueryRow(userId, guildId).Scan(&userNum)

// 	// 여러 행 : Query
// 	// var userNum int
// 	// res, err := stmt.Query(userId, guildId)
// 	// for res.Next() {
// 	// 	res.Scan(&userNum)
// 	// }
// 	stmt.Close()
// 	return userNum, err
// }

// Todo 삭제 - 사유 : 빠따 -> 성공 회수
// func SelectBbadda(userNum int) (int, error) {
// 	db := getConnection()
// 	err := db.Ping()
// 	var bbadda int64
// 	if err == nil {
// 		sql := `SELECT bbadda
// 		FROM user WHERE userNum = ?`
// 		stmt, _ := db.Prepare(sql)
// 		err = stmt.QueryRow(userNum).Scan(&bbadda)
// 		stmt.Close()
// 	}
// 	defer db.Close()
// 	db.Close()
// 	return int(bbadda), err
// }

// Todo 삭제 - 사유 : 빠따 -> 성공 회수
// func SelectOverBbadaa(guildId string, bbaddaLimit int) ([]ds.User, error) {
// 	db := getConnection()
// 	err := db.Ping()
// 	var user ds.User
// 	var users []ds.User
// 	if err == nil {
// 		sql := `select userNum, userId, guildId, userName, bbadda, userType from user where guildId = ? and bbadda > ?`
// 		stmt, _ := db.Prepare(sql)
// 		res, _ := stmt.Query(guildId, bbaddaLimit)

// 		for res.Next() {
// 			res.Scan(&user.UserNum, &user.UserId, &user.GuildId, &user.UserName, &user.Bbadda, &user.UserType)
// 			users = append(users, user)
// 		}
// 		stmt.Close()
// 	}
// 	defer db.Close()
// 	return users, err
// }

// Todo 삭제 - 사유 : 빠따 -> 성공 회수
// 빠따 카운트 + 1
// func UpdateBbaddaByUserNum(userNum int) (*int64, error) {
// 	db := getConnection()
// 	err := db.Ping()
// 	var no int64
// 	if err == nil {
// 		sql := `UPDATE user SET bbadda = bbadda+1
// 		WHERE userNum = ?`

// 		stmt, _ := db.Prepare(sql)
// 		res, _ := stmt.Exec(userNum)

// 		no, err = res.RowsAffected()
// 		stmt.Close()
// 	}
// 	defer db.Close()
// 	return &no, err
// }
