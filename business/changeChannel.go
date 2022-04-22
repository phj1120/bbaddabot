/*
작성자 : 박현준
작성일 : 2022.03.19.

수정자 : 박현준
수정일 : 2022.04.22.

파일 설명
채널 이동 시 수행될 기능이 있는 파일
*/

package business

import (
	ds "bbaddabot/datastruct"
	ps "bbaddabot/persistence"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// msg 를 받아 discord 메시지와 log 메시지 한 번에 전송
func sendMessageAndPrintLog(s *discordgo.Session, channelIds []string, msg string) {
	if channelIds == nil {
		s.ChannelMessageSend("952057033476177920", msg)
	} else {
		for _, channelId := range channelIds {
			s.ChannelMessageSend(channelId, msg)
		}
	}
	fmt.Println(msg)
}

// 사용자의 채널이 변경 될 경우 상황에 맞는 기능 수행
func ChangeChannel(s *discordgo.Session, v discordgo.VoiceStateUpdate) {

	// 모든 메시지가 전송될 채널 ( 채널 이동 목록 전체 )
	allLogChannelIds := ps.SelectChannelIdByChannelTypeAndGuildId("로그", v.GuildID)
	// 주요 메시지가 전송될 채널 ( 10분 이상의 공부 시간 )
	studyLogChannelIds := ps.SelectChannelIdByChannelTypeAndGuildId("기록", v.GuildID)

	var subMsg string
	var msg string
	var totalStudyTime int
	histoty := ds.History{}

	userNum, err := ps.SelectUserNumByUserIdAndGuildId(v.UserID, v.GuildID)
	ps.SelectUserNumByUserIdAndGuildId(v.UserID, v.GuildID)

	// 유저가 없는 경우 유저 추가
	if err != nil {
		user := ds.User{}
		user.UserId = v.UserID
		user.GuildId = v.GuildID
		userTmp, _ := s.User(v.UserID)
		user.UserName = userTmp.Username
		ps.UserSave(user)
		user, _ = ps.UserFindByUserNum(histoty.UserNum)
		userNum, _ = ps.SelectUserNumByUserIdAndGuildId(v.UserID, v.GuildID)
	}

	histoty.UserNum = userNum
	user, _ := ps.UserFindByUserNum(histoty.UserNum)
	userName := user.UserName

	// 채널간 이동이 발생한 경우
	if v.BeforeUpdate != nil && v.VoiceState != nil {
		// 동일 채널일 경우 처리할 것 없음
		if v.BeforeUpdate.ChannelID == v.VoiceState.ChannelID {
			sendMessageAndPrintLog(s, allLogChannelIds, "!change user state in same channel")
			return
		}
		// 이전 채널이 있는 경우 - 퇴장, 타채널로 이동
		if v.BeforeUpdate.ChannelID != "" {
			histoty.HistoryType = ps.SelectChannelTypeById(v.BeforeUpdate.ChannelID)
			histoty.BeforeChannelId = v.BeforeUpdate.ChannelID
			beforeChannelName := ps.SelectChannelNameById(histoty.BeforeChannelId)

			// 이후 채널이 없는 경우 - 퇴장
			if v.ChannelID == "" {
				subMsg = fmt.Sprintf("%s%s", " 종료 : ", beforeChannelName)
				// 이후 채널이 있는 경우 - 타채널로 이동
			} else {
				histoty.AfterChannelId = v.ChannelID
				afterChannelName := ps.SelectChannelNameById(histoty.AfterChannelId)
				subMsg = fmt.Sprintf("%s%s%s%s", " 이동 : ", beforeChannelName, " -> ", afterChannelName)
			}
		}

		// 이동 기록 삽입
		_, err := ps.InsertHistory(histoty)
		if err != nil {
			msg = fmt.Sprintf("%s%#v", "!insert user error - 이동, 종료", histoty)
			sendMessageAndPrintLog(s, allLogChannelIds, msg)
			return
		}

		spentMinute := ps.SelectMinuteSpentByUserNum(userNum)
		// 공부 기록인 경우 총합 공부 시간 갱신
		if histoty.HistoryType == "공부" {
			totalStudyTime, err = ps.SelectStudyTotalTodayByUserNum(histoty.UserNum)
			if err != nil {
				// 당일에 처음 기록 하는 경우
				ps.InsertNewStudyTotal(histoty.UserNum, spentMinute)
			} else {
				ps.UpdateStudyTimeByUserNumAndStudyTime(histoty.UserNum, spentMinute)
			}
			msg = fmt.Sprintf("| %s | %s %s | %s / %s | %s |", time.Now().Format("20060102 15:04:05"), userName, histoty.HistoryType, minuteToHour(spentMinute), minuteToHour(totalStudyTime+spentMinute), subMsg)
			if spentMinute > 10 {
				subMsg = fmt.Sprintf("[ %s ] %s : %s / %s", time.Now().Format("20060102 15:04"), userName, minuteToHour(spentMinute), minuteToHour(totalStudyTime+spentMinute))
				sendMessageAndPrintLog(s, studyLogChannelIds, subMsg)
			}
		} else {
			msg = fmt.Sprintf("| %s | %s %s | %s | %s |", time.Now().Format("20060102 15:04:05"), userName, histoty.HistoryType, minuteToHour(spentMinute), subMsg)
		}
	}

	// 이전 채널이 없는 경우 - 입장
	if v.BeforeUpdate == nil {
		histoty.AfterChannelId = v.ChannelID
		histoty.HistoryType = "start"
		afterChannelName := ps.SelectChannelNameById(v.ChannelID)
		subMsg = fmt.Sprintf(" 입장 : %s", afterChannelName)

		// 이동 기록 삽입
		_, err := ps.InsertHistory(histoty)
		if err != nil {
			msg = fmt.Sprintf("%s%#v", "!insert user error - 입장", histoty)
			sendMessageAndPrintLog(s, allLogChannelIds, msg)
		}
		msg = fmt.Sprintf("| %s | %s | %s |", time.Now().Format("20060102 15:04:05"), userName, subMsg)
	}
	sendMessageAndPrintLog(s, allLogChannelIds, msg)
}
