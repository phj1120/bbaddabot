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
	// bbadda, _ := ps.SelectBbadda(userNum)
	user, _ := ps.UserFindByUserNum(userNum)
	studyTime, _ := ps.SelectStudyTotalTodayByUserNum(userNum)

	var msg string
	var subMsg string
	if request == "!기록" || request == "!ㄱㄹ" {
		userList, _ := ps.UserListFindByGuildId(m.GuildID)
		msg = fmt.Sprintf("[%s]\n", time.Now().Format("20060102 15:04"))
		for _, user = range userList {
			// bbadda, _ = ps.SelectBbadda(user.UserNum)
			studyTime, _ = ps.SelectStudyTotalTodayByUserNum(user.UserNum)
			subMsg = fmt.Sprintf("%s : %s\n", user.UserName, minuteToHour(studyTime))
			// subMsg = fmt.Sprintf("%s : %s / %d 개\n", user.UserName, minuteToHour(studyTime), bbadda)
			msg += subMsg
		}
	}

	if request == "!공부시간" || request == "!ㄱㅄㄱ" {
		msg = fmt.Sprintf("[%s] %s : %s", time.Now().Format("20060102 15:04"), user.UserName, minuteToHour(studyTime))
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

	// 강퇴 기능 추가 중
	// slice := strings.Split(request, ".")
	// fmt.Println(slice)
	// if slice[0] == "!강퇴" {
	// 	user, _ := persistence.SelectUserByUserNum(userNum)
	// 	if slice[1] == user.UserName {
	// 		fmt.Println(m.GuildID)
	// 		fmt.Println(user.UserId)
	// 		err := s.GuildMemberDeleteWithReason(m.GuildID, user.UserId, "빠따 초과")
	// 		fmt.Println(err)
	// 	}
	// }

	// Todo - 빠따 삭제로 관련 기능 삭제
	// if request == "!빠따" {
	// 	msg = fmt.Sprintf("[%s] %s : %d 개", time.Now().Format("20060102 15:04"), user.UserName, bbadda)
	// }
	//
	// if request == "!정산" {
	// 	userList, _ := ps.SelectUserList(m.GuildID)
	//
	// 	msg = fmt.Sprintf("[%s]\n", time.Now().Format("20060102 15:04"))
	// 	for _, user = range userList {
	// 		bbadda, _ = ps.SelectBbadda(user.UserNum)
	// 		studyTime, _ = ps.SelectStudyTotalTodayByUserNum(user.UserNum)
	// 		if studyTime <= 180 {
	// 			ps.UpdateBbaddaByUserNum(user.UserNum)
	// 			bbadda, _ = ps.SelectBbadda(user.UserNum)
	// 		}
	// 		subMsg = fmt.Sprintf("%s : %s / %d 개\n", user.UserName, minuteToHour(studyTime), bbadda)
	// 		msg += subMsg
	// 	}
	// }
	s.ChannelMessageSend(m.ChannelID, msg)

}

func minuteToHour(minute int) string {
	var m int
	var h int
	var msg string

	if minute >= 60 {
		h = minute / 60
		m = minute - h*60
		msg = fmt.Sprintf("%d 시간 %d 분", h, m)
	} else {
		msg = fmt.Sprintf("%d 분", minute)
	}
	return msg
}
