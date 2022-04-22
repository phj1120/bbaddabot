/*
작성자 : 박현준
작성일 : 2022.03.19.

수정자 : 김민우
수정일 : 2022.04.22.

파일 설명
사용자의 메시지를 받아 답해주는 챗 봇 기능이 있는 파일
*/

package business

import (
	_ "bbaddabot/datastruct"
	ps "bbaddabot/persistence"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Chatbot(s *discordgo.Session, m *discordgo.MessageCreate) {
	request := m.Content

	userNum, _ := ps.SelectUserNumByUserIdAndGuildId(m.Author.ID, m.GuildID)
	user, _ := ps.UserFindByUserNum(userNum)

	var msg string
	var subMsg string
	var studyTime int

	if request == "!기록" || request == "!ㄱㄹ" {
		userList, _ := ps.UserListFindByGuildId(m.GuildID)
		msg = fmt.Sprintf("[%s]\n", time.Now().Format("20060102 15:04"))
		for _, user = range userList {
			studyTime, _ = ps.SelectStudyTotalTodayByUserIdAndGuildId(user.UserId, user.GuildId)
			subMsg = fmt.Sprintf("%s : %s\n", user.UserName, minuteToHour(studyTime))
			msg += subMsg
		}
	}

	if request == "!일" {
		studyTime, _ = ps.SelectStudyTotalTodayByUserIdAndGuildId(m.Author.ID, m.GuildID)
		msg = fmt.Sprintf("[%s] %s : %s", time.Now().Format("2006년 01월 02일"), user.UserName, minuteToHour(studyTime))
	}

	if request == "!주" {
		studyTime, _ = ps.SelectStudyTotalWeekByUserIdAndGuildId(m.Author.ID, m.GuildID)

		now := time.Now()
		firstWeekDay := time.Date(now.Year(), now.Month(), 1, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
		// 현재 주차 = ( 현재 요일 + 1일 요일 숫자 - 1) / 7 + 1
		weekNumber := (now.Day()+int(firstWeekDay.Weekday())-1)/7 + 1
		msg = fmt.Sprintf("[%s %d주차] %s: %s", time.Now().Format("2006년 01월"), weekNumber, user.UserName, minuteToHour(studyTime))
	}

	if request == "!달" {
		studyTime, _ = ps.SelectStudyTotalMonthByUserIdAndGuildId(m.Author.ID, m.GuildID)
		msg = fmt.Sprintf("[%s] %s : %s", time.Now().Format("2006년 01월"), user.UserName, minuteToHour(studyTime))
	}

	// Server setting command
	if strings.HasPrefix(request, "!설정") {
		setting := strings.Split(request, " ")

		if len(setting) <= 1 {
			msg = "---------설정명령어---------\n"
			msg += "!설정 채팅채널 : 채널 현재 설정상태 출력\n"
			msg += "!설정 채팅채널 [채널타입] : 특정 채널 타입으로 설정"
		} else {
			switch {
			// Set Channel Type
			case setting[1] == "채팅채널":

				// set channel type If there's flag after "채팅채널"
				if len(setting) == 3 {
					channelType := setting[2]

					msg = "---------채널 설정 완료---------\n"
					msg += "이전 채널 이름 : " + ps.SelectChannelNameById(m.ChannelID)
					msg += " 이전 채널 설정 : " + ps.SelectChannelTypeById(m.ChannelID)
					ps.UpdateChannelType(m.ChannelID, channelType) // Set channel type
				}

				msg += "\n현재 채널 이름 : " + ps.SelectChannelNameById(m.ChannelID)
				msg += " 현재 채널 설정 : " + ps.SelectChannelTypeById(m.ChannelID)

				// Guide when no info in DB
				/*
					if ps.SelectChannelNameById(m.ChannelID) == "" || ps.SelectChannelTypeById(m.ChannelID) == "" {
						msg += "\n채널 설정 정보가 없습니다. 채널을 설정해주십시오."
						msg += "\n명령어: !설정 채팅채널 [채널타입]"
					} else {
						// show current channel type
						msg += "\n현재 채널 이름 : " + ps.SelectChannelNameById(m.ChannelID)
						msg += " 현재 채널 설정 : " + ps.SelectChannelTypeById(m.ChannelID)
					}*/

			// Set voice channel type
			case setting[1] == "음성채널":
				msg = "기능준비중"

			// Show help message if command is none of listed
			default:
				msg = "---------설정명령어---------\n"
				msg += "!설정 채팅채널 : 채널 현재 설정상태 출력\n"
				msg += "!설정 채팅채널 [채널타입] : 특정 채널 타입으로 설정"
			}

		}

	}
	s.ChannelMessageSend(m.ChannelID, msg)
}

func minuteToHour(minute int) string {
	var d int
	var m int
	var h int
	var msg string
	println(minute)
	if minute > 60*24 {
		d = minute / (24 * 60)
		h = (minute - d*24*60) / 60
		m = minute - h*60 - d*24*60
		msg = fmt.Sprintf("%d 일 %d 시간 %d 분", d, h, m)
	} else if minute >= 60 {
		h = minute / 60
		m = minute - h*60
		msg = fmt.Sprintf("%d 시간 %d 분", h, m)
	} else {
		msg = fmt.Sprintf("%d 분", minute)
	}
	return msg
}
