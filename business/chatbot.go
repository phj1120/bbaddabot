package business

import (
	_ "bbaddabot/datastruct"
	ps "bbaddabot/persistence"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Chatbot(s *discordgo.Session, m *discordgo.MessageCreate) {

	request := m.Content

	userNum, _ := ps.SelectUserNumByUserIdAndGuildId(m.Author.ID, m.GuildID)
	bbadda, _ := ps.SelectBbadda(userNum)
	user, _ := ps.SelectUserByUserNum(userNum)
	studyTime, _ := ps.SelectStudyTotalTodayByUserNum(userNum)

	var msg string
	var subMsg string
	if request == "!기록" || request == "!ㄱㄹ" {
		userList, _ := ps.SelectUserList(m.GuildID)

		msg = fmt.Sprintf("[%s]\n", time.Now().Format("20060102 15:04"))
		for _, user = range userList {
			bbadda, _ = ps.SelectBbadda(user.UserNum)
			studyTime, _ = ps.SelectStudyTotalTodayByUserNum(user.UserNum)
			subMsg = fmt.Sprintf("%s : %s / %d 개\n", user.UserName, minuteToHour(studyTime), bbadda)
			msg += subMsg
		}
	}

	if request == "!빠따" {
		msg = fmt.Sprintf("[%s] %s : %d 개", time.Now().Format("20060102 15:04"), user.UserName, bbadda)
	}

	if request == "!공부시간" || request == "!ㄱㅄㄱ" {
		msg = fmt.Sprintf("[%s] %s : %s", time.Now().Format("20060102 15:04"), user.UserName, minuteToHour(studyTime))
	}

	if request == "!정산" {
		userList, _ := ps.SelectUserList(m.GuildID)

		msg = fmt.Sprintf("[%s]\n", time.Now().Format("20060102 15:04"))
		for _, user = range userList {
			bbadda, _ = ps.SelectBbadda(user.UserNum)
			studyTime, _ = ps.SelectStudyTotalTodayByUserNum(user.UserNum)
			if studyTime <= 180 {
				ps.UpdateBbaddaByUserNum(user.UserNum)
				bbadda, _ = ps.SelectBbadda(user.UserNum)
			}
			subMsg = fmt.Sprintf("%s : %s / %d 개\n", user.UserName, minuteToHour(studyTime), bbadda)
			msg += subMsg
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
