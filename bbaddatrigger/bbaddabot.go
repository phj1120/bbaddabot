package main

import (
	ps "bbaddabot/persistence"
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

var offSignal bool = false

func main() {
	token := os.Getenv("BBADDABOTTOKEN")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(swing)

	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	for offSignal == false {
	}

	dg.Close()
}

func swing(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("swing")
	guildList, _ := ps.SelectGuildIdList()
	for _, guildId := range guildList {
		sendTotal(s, r, guildId)
		sendBbadda(s, r, guildId)
	}
	offSignal = true
}

func sendBbadda(s *discordgo.Session, r *discordgo.Ready, guildId string) {
	var outCnt int = 0
	var msg string
	userList, _ := ps.SelectUserByGuildId(guildId)

	for _, user := range userList {
		if user.Bbadda > 4 { // 추후 개별 목표 DB 에서 가져오도록 추가 할 것
			err := s.GuildMemberDeleteWithReason(guildId, user.UserId, "빠따 초과")
			msg += fmt.Sprintf("빠따 초과 : %+v\n", user)
			outCnt += 1
			if err != nil {
				fmt.Println("GuildMemberDeleteWithReason err")
				fmt.Println(err)
			} else {
				ps.DeleteUser(user.UserId)
			}
		}
	}
	msg += fmt.Sprintf("%d 명 아웃\n", outCnt)

	notiyChannel := ps.SelectNotiyChannel(guildId)
	if notiyChannel != "" {
		s.ChannelMessageSend(notiyChannel, msg)
	}
}

func sendTotal(s *discordgo.Session, r *discordgo.Ready, guildId string) {
	fmt.Println("sendTotal")
	userList, _ := ps.SelectUserByGuildId(guildId)

	var msg string
	var subMsg string
	msg = fmt.Sprintf("[%s]\n", time.Now().Format("20060102 15:04"))
	for _, user := range userList {
		bbadda, _ := ps.SelectBbadda(user.UserNum)
		studyTime, _ := ps.SelectStudyTotalTodayByUserNum(user.UserNum)
		subMsg = fmt.Sprintf("%s : %s / %d 개\n", user.UserName, minuteToHour(studyTime), bbadda)
		msg += subMsg
	}
	notiyChannel := ps.SelectNotiyChannel(guildId)
	if notiyChannel != "" {
		s.ChannelMessageSend(notiyChannel, msg)
	}
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
