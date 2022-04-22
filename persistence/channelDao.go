/*
작성자 : 박현준
작성일 : 2022.03.19.

수정자 : 박현준
수정일 : 2022.04.22.

파일 설명
Channel(No, GuildId, ChannelId, ChannelName, ChannelType) 테이블 매핑
*/

package persistence

import (
	ds "bbaddabot/datastruct"

	_ "github.com/go-sql-driver/mysql"
)

// 채널 추가
func InsertChannel(c ds.Channel) int {
	db := getConnection()
	err := db.Ping()
	var id int64
	if err == nil {
		sql := `INSERT INTO  channel (guildId, channelId, channelName, channelType) 
		VALUES(?, ?, ?, ?);`

		stmt, _ := db.Prepare(sql)
		res, _ := stmt.Exec(c.GuildId, c.ChannelId, c.ChannelName, c.ChannelType)
		id, _ = res.RowsAffected()
		stmt.Close()
	}
	return int(id)
}

// 채널 아이디로 채널 이름 조회
func SelectChannelNameById(chnnelId string) string {
	db := getConnection()
	err := db.Ping()
	var channelName string
	if err == nil {
		sql := `SELECT channelName
				FROM channel
				WHERE channelId = ?`
		stmt, _ := db.Prepare(sql)

		stmt.QueryRow(chnnelId).Scan(&channelName)
		stmt.Close()
	}
	return channelName
}

// 채널 아이디로 채널 타입 조회
func SelectChannelTypeById(channelId string) string {
	db := getConnection()
	err := db.Ping()
	var channelType string
	if err == nil {
		sql := `SELECT channelType
		FROM channel
		WHERE channelId = ?`
		stmt, _ := db.Prepare(sql)

		stmt.QueryRow(channelId).Scan(&channelType)
		stmt.Close()
	}
	return channelType
}

// 채널 길드ID 와 채널 타입으로 채널 아이디 조회
func SelectChannelIdByChannelTypeAndGuildId(channelType string, guildId string) []string {
	db := getConnection()
	err := db.Ping()
	var channelIds []string
	if err == nil {
		sql := `SELECT channelId
		FROM channel
		WHERE channelType = ? and guildId = ?`
		stmt, _ := db.Prepare(sql)
		res, _ := stmt.Query(channelType, guildId)

		var channelId string

		for res.Next() {
			res.Scan(&channelId)
			channelIds = append(channelIds, channelId)
		}
		stmt.Close()
	}
	return channelIds
}

// 채널 타입 변경
func UpdateChannelType(channelId string, channelType string) int {
	db := getConnection()
	err := db.Ping()
	var cnt int64
	if err == nil {
		sql := `UPDATE channel SET channelType = ?
		WHERE channelId = ?`

		stmt, _ := db.Prepare(sql)
		res, _ := stmt.Exec(channelType, channelId)
		cnt, _ = res.RowsAffected()
		stmt.Close()
	}
	return int(cnt)
}
