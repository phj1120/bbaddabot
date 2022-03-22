package business

import (
	ds "bbaddabot/datastruct"
	ps "bbaddabot/persistence"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func ChangeChannel(s *discordgo.Session, v discordgo.VoiceStateUpdate) string {
	// fmt.Println("Change Channel")
	var subMsg string
	var msg string
	h := ds.History{}

	// 유저가 없는 경우 유저 추가
	userNum, err := ps.SelectUserNumByUserIdAndGuildId(v.UserID, v.GuildID)
	ps.SelectUserNumByUserIdAndGuildId(v.UserID, v.GuildID)

	if err != nil {
		user := ds.User{}
		user.UserId = v.UserID
		user.GuildId = v.GuildID
		userTmp, _ := s.User(v.UserID)
		user.UserName = userTmp.Username
		ps.InsertUser(user)
		user, _ = ps.SelectUserByUserNum(h.UserNum)
		userNum, _ = ps.SelectUserNumByUserIdAndGuildId(v.UserID, v.GuildID)
	}
	h.UserNum = userNum
	// UserNum 으로 User 조회
	user, err := ps.SelectUserByUserNum(h.UserNum)
	userName := user.UserName

	// 채널간 이동이 발생한 경우
	if v.BeforeUpdate != nil && v.VoiceState != nil {
		// 동일 채널일 경우 처리할 것 없음
		if v.BeforeUpdate.ChannelID == v.VoiceState.ChannelID {
			return "!change user state in same channel"
		}
		// 이전 채널이 있는 경우 - 퇴장, 타채널로 이동
		if v.BeforeUpdate.ChannelID != "" {
			h.HistoryType = ps.SelectChannelTypeById(v.BeforeUpdate.ChannelID)
			h.BeforeChannelId = v.BeforeUpdate.ChannelID
			beforeChannelName := ps.SelectChannelNameById(h.BeforeChannelId)

			// 이후 채널이 없는 경우 - 퇴장
			if v.ChannelID == "" {
				subMsg = fmt.Sprintf("%s%s", " / 종료 : ", beforeChannelName)
				// 이후 채널이 있는 경우 - 타채널로 이동
			} else {
				h.AfterChannelId = v.ChannelID
				afterChannelName := ps.SelectChannelNameById(h.AfterChannelId)
				subMsg = fmt.Sprintf("%s%s%s%s", " / 이동 : ", beforeChannelName, " -> ", afterChannelName)
			}
		}
		// 이동 기록 삽입
		_, err := ps.InsertHistory(h)
		if err != nil {
			msg = fmt.Sprintf("%s%#v", "!insert user error - 이동, 종료", h)
			fmt.Println(msg)
			return msg
		}

		// 채널에 있었던 시간(분) 계산
		spentMinute := ps.SelectMinuteSpentChannel2(userNum)

		// 공부 기록인 경우 총합 공부 시간 갱신
		if h.HistoryType == "공부" {
			// 총 공부 시간
			_, err := ps.SelectStudyTotalTodayByUserNum(h.UserNum)
			if err != nil {
				// 당일에 처음 기록 하는 경우
				ps.InsertNewStudyTotal(h.UserNum, spentMinute)
			} else {
				ps.UpdateStudyTimeByUserNumAndStudyTime(h.UserNum, spentMinute)
			}
		}
		msg = fmt.Sprintf("%s%s%s%s%d%s%s%s", time.Now().Format("20060102 15:04:05"), " ", userName, " ", spentMinute, " 분 ", h.HistoryType, subMsg)
	}

	// 이전 채널이 없는 경우 - 입장
	if v.BeforeUpdate == nil {
		h.AfterChannelId = v.ChannelID
		h.HistoryType = "start"
		afterChannelName := ps.SelectChannelNameById(v.ChannelID)
		subMsg = fmt.Sprintf("%s%s", " / 입장 : ", afterChannelName)

		// 이동 기록 삽입
		_, err := ps.InsertHistory(h)
		if err != nil {
			msg = fmt.Sprintf("%s%#v", "!insert user error - 입장", h)
			fmt.Println(msg)
		}
		msg = fmt.Sprintf("%s%s%s%s", time.Now().Format("20060102 15:04:05"), " ", userName, subMsg)
	}

	return msg
}
