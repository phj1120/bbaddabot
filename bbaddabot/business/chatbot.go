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
	if request == "!기록" {
		userList, _ := ps.SelectUserList(m.GuildID)
		msg = fmt.Sprintf("[%s]\n", time.Now().Format("20060102 15:04"))
		for _, user = range userList {
			bbadda, _ = ps.SelectBbadda(user.UserNum)
			studyTime, _ = ps.SelectStudyTotalTodayByUserNum(user.UserNum)
			// subMsg = fmt.Sprintf("[%s] %s : 총 %d 분 / %d 개\n", time.Now().Format("20060102 15:04"), user.UserName, studyTime, bbadda)
			subMsg = fmt.Sprintf("%s : 총 %d 분 / %d 개\n", user.UserName, studyTime, bbadda)
			msg += subMsg
		}
		s.ChannelMessageSend(m.ChannelID, msg)
	}

	if request == "!빠따" {
		msg = fmt.Sprintf("[%s] %s : 총 %d 개", time.Now().Format("20060102 15:04"), user.UserName, bbadda)
		s.ChannelMessageSend(m.ChannelID, msg)
	}

	if request == "!공부시간" {
		msg = fmt.Sprintf("[%s] %s : 총 %d 분", time.Now().Format("20060102 15:04"), user.UserName, studyTime)
		s.ChannelMessageSend(m.ChannelID, msg)
	}
}
