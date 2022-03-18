package business

import (
	ds "bbaddabot/datastruct"
	ps "bbaddabot/persistence"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func ChangeChannel(v discordgo.VoiceStateUpdate) string {
	fmt.Println("Change Channel")
	var subMsg string
	var msg string
	h := ds.History{}
	h.UserNum, _ = ps.SelectUserNumByUserIdAndGuildId(v.UserID, v.GuildID)
	user, _ := ps.SelectUserByUserNum(h.UserNum)
	userName := user.UserName

	// 채널간 이동이 발생한 경우
	if v.BeforeUpdate != nil && v.VoiceState != nil {
		// 동일 채널일 경우 처리할 것 없음
		if v.BeforeUpdate.ChannelID == v.VoiceState.ChannelID {
			return "err"
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
		LastUpdateNo, err := ps.InsertHistory(h)
		if err != nil {
			return "err"
		}

		// 채널에 있었던 시간(분) 계산
		spentMinute := ps.SelectMinuteSpentChannel(LastUpdateNo)

		// 공부 기록인 경우 총합 공부 시간 갱신
		if h.HistoryType == "study" {
			// 총 공부 시간
			_, err := ps.SelectStudyTotalTodayByUserNum(h.UserNum)
			if err != nil {
				// 당일에 처음 기록 하는 경우
				ps.InsertNewStudyTotal(h.UserNum, spentMinute)
			} else {
				ps.UpdateStudyTimeByUserNumAndStudyTime(h.UserNum, spentMinute)
			}
		}
		msg = fmt.Sprintf("%s%s%s%s%d%s%s%s", time.Now(), " ", userName, " ", spentMinute, " 분 ", h.HistoryType, subMsg)
	}

	// 이전 채널이 없는 경우 - 입장
	if v.BeforeUpdate == nil {
		afterChannelName := ps.SelectChannelNameById(v.ChannelID)
		subMsg = fmt.Sprintf("%s%s", " / 입장 : ", afterChannelName)

		// 이동 기록 삽입
		_, err := ps.InsertHistory(h)
		if err != nil {
			return "err"
		}
		msg = fmt.Sprintf("%s%s%s%s", time.Now(), " ", userName, subMsg)
	}

	return msg
}
