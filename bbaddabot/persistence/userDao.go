package persistence

import (
	ds "bbaddabot/datastruct"

	_ "github.com/go-sql-driver/mysql"
)

// userNum, userId, guildId, userName, bbadda, type
func InsertUser(u ds.User) (*int64, error) {
	db := dbConn()
	sql := `INSERT INTO  user (userId, guildId, userName, bbadda, userType) 
	VALUES(?, ?, ?, 0, 'user')`

	stmt, _ := db.Prepare(sql)
	res, _ := stmt.Exec(u.UserId, u.GuildId, u.UserName)
	id, err := res.LastInsertId()

	return &id, err
}

// User 조회
func SelectUserByUserNum(userNum int) (ds.User, error) {
	db := dbConn()
	sql := `SELECT userNum, userId, guildId, userName, bbadda, userType
	FROM user WHERE userNum=?`
	stmt, _ := db.Prepare(sql)

	var user ds.User
	err := stmt.QueryRow(userNum).Scan(&user.UserNum, &user.UserId, &user.GuildId, &user.UserName, &user.Bbadda, &user.UserType)
	return user, err
}

// UserNum 조회
func SelectUserNumByUserIdAndGuildId(userId string, guildId string) (int, error) {
	db := dbConn()
	sql := `SELECT userNum
	FROM user WHERE userId=? AND guildId=?`
	stmt, _ := db.Prepare(sql)

	// 행 하나 : QueryRow
	var userNum int
	err := stmt.QueryRow(userId, guildId).Scan(&userNum)

	// 여러 행 : Query
	// var userNum int
	// res, err := stmt.Query(userId, guildId)
	// for res.Next() {
	// 	res.Scan(&userNum)
	// }
	return userNum, err
}

// 유저 타입
func UpdateUserType(u ds.User) (*int64, error) {
	db := dbConn()
	sql := `UPDATE user SET type = ?
	WHERE userNum = ?`

	stmt, _ := db.Prepare(sql)
	res, _ := stmt.Exec(u.UserType, u.UserId)

	// auto_increment 반환
	no, err := res.LastInsertId()

	return &no, err
}

// 빠따 카운트 + 1
func UpdateUserBbaddaByUserId(userId string) (*int64, error) {
	db := dbConn()
	sql := `UPDATE user SET bbadda = bbadda+1
	WHERE userId = ?`

	stmt, _ := db.Prepare(sql)
	res, _ := stmt.Exec(userId)

	// auto_increment 반환
	no, err := res.RowsAffected()

	return &no, err
}

// 빠따 카운트 > n 인 사용자 삭제
// func DeleteUserOverBbadda(n int) (*int64, error) {
// 	db := dbConn()
// 	sql := `DELETE FROM user
// 	WHERE bbadda => ?`

// 	stmt, _ := db.Prepare(sql)
// 	res, _ := stmt.Exec(n)
// 	// 반영된 row 수 반환
// 	cnt, err := res.RowsAffected()

// 	return &cnt, err
// }
