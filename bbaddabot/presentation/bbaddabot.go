package presentation

import (
	"bbaddabot/business"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

func PresentationTest() {
	fmt.Println("PresentationTest")
	business.BusinessTest()
}

func Bbaddabot() {
	token := os.Getenv("BBADDABOTTOKEN")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.AddHandler(myVociceStatusUpdate)

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Airhorn is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func myVociceStatusUpdate(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	if v.VoiceState == nil {
		return
	}

	userName := getUserName(s, v.VoiceState.UserID)
	nowTime := time.Now().Format("20060102 15:04:05")
	msg := "None"

	// 지금 이거 너무 비효율 적임 일단 이렇게 짜고 로직 수정하기
	// 채널 이름 읽고 그러는거  오래 걸린다.
	// 채널 여기서 안읽고 하고 DB 에서 읽기?
	if v.BeforeUpdate != nil && v.VoiceState.ChannelID == "" {
		beforeChannelName := getChannelName(s, v.BeforeUpdate.ChannelID)
		msg = fmt.Sprintf("%s%s%s%s%s%s", nowTime, " ", userName, " ", beforeChannelName, " 종료")
		fmt.Println(msg)
	} else if v.BeforeUpdate != nil && v.VoiceState != nil {
		beforeChannelName := getChannelName(s, v.BeforeUpdate.ChannelID)
		afterChannelName := getChannelName(s, v.VoiceState.ChannelID)
		msg = fmt.Sprintf("%s%s%s%s%s%s%s", nowTime, " ", userName, " 이동 : ", beforeChannelName, " -> ", afterChannelName)
		fmt.Println(msg)
	} else if v.BeforeUpdate == nil && v.VoiceState.ChannelID != "" {
		afterChannelName := getChannelName(s, v.VoiceState.ChannelID)
		msg = fmt.Sprintf("%s%s%s%s%s%s", nowTime, " ", userName, " ", afterChannelName, " 입장")
		fmt.Println(msg)
	}
	s.ChannelMessageSend("952057033476177920", msg)
}

func getChannelName(s *discordgo.Session, channelID string) string {
	channel, _ := s.Channel(channelID)
	return channel.Name
}

func getUserName(s *discordgo.Session, userID string) string {
	user, _ := s.User(userID)
	return user.Username
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	currntChannel, _ := s.Channel(m.ChannelID)
	currntChannelName := currntChannel.Name

	msg := fmt.Sprintf("%s %s %s %s", m.Timestamp.Format("20060102 15:04:05"), currntChannelName, m.Author.Username, m.Content)

	// 채팅 로그 채널에 메시지 전송
	s.ChannelMessageSend("952040735090294804", msg)
	// 출력
	fmt.Println(msg)
}
