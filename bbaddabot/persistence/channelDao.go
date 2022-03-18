package persistence

import (
	ds "bbaddabot/datastruct"

	_ "github.com/go-sql-driver/mysql"
)

// guildId, channelId, channelName, channelType

// 채널 추가
func InsertChannel(c ds.Channel) int {
	db := dbConn()
	sql := `INSERT INTO  channel (guildId, channelId, channelName, channelType) 
	VALUES(?, ?, ?, ?);`

	stmt, _ := db.Prepare(sql)
	res, _ := stmt.Exec(c.GuildId, c.ChannelId, c.ChannelName, c.ChannelType)
	id, _ := res.RowsAffected()

	return int(id)
}

// 채널 아이디로 채널 이름 조회
func SelectChannelNameById(chnnelId string) string {
	db := dbConn()
	sql := `SELECT channelName
			FROM channel
			WHERE channelId = ?`
	stmt, _ := db.Prepare(sql)

	var channelName string
	stmt.QueryRow(chnnelId).Scan(&channelName)
	return channelName
}

// 채널 아이디로 채널 이름 조회
func SelectChannelTypeById(channelId string) string {
	db := dbConn()
	sql := `SELECT channelType
			FROM channel
			WHERE channelId = ?`
	stmt, _ := db.Prepare(sql)

	var channelType string
	stmt.QueryRow(channelId).Scan(&channelType)
	return channelType
}

// 채널 타입 변경
func UpdateChannelType(channelId string, channelType string) int {
	db := dbConn()
	sql := `UPDATE channel SET channelType = ?
	WHERE channelId = ?`

	stmt, _ := db.Prepare(sql)
	res, _ := stmt.Exec(channelType, channelId)
	cnt, _ := res.RowsAffected()

	return int(cnt)
}
