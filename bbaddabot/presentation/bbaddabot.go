package presentation

import (
	// "bbaddabot/business"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

func PresentationTest() {
	// fmt.Println("PresentationTest")
	// business.BusinessTest()
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
	// dg.AddHandler(myVociceStatusUpdate)
	// dg.AddHandler(threadUpdate)

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("BBADDABOT is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func log(s *discordgo.Session, msg string) {
	fmt.Println(msg)
	s.ChannelMessageSend("954047325448335410", msg)
}

// 채널 편집 시 이벤트 발생
func channelUpdate(s *discordgo.Session, c *discordgo.ChannelUpdate) {
	msg := c.ID
	log(s, msg)
}

// 스레드.. 텍스트 메시지 관련된것 같은데 잘 모르겠네
func threadUpdate(s *discordgo.Session, t *discordgo.ThreadUpdate) {
	msg := t.Channel.ID
	log(s, msg)
}

// 재접속일 줄 알았는데 내가 원하는 재접속은 아니었음
func presenceUpdate(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	msg := p.User
	log(s, msg.ID)
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

	// 강퇴 성공
	strs := m.Content
	slice := strings.Split(strs, "=")

	if slice[0] == "!강퇴" {
		s.GuildMemberDeleteWithReason(m.GuildID, slice[1], "강퇴")
	}

	currntChannel, _ := s.Channel(m.ChannelID)
	currntChannelName := currntChannel.Name

	// 제자리 이동 가능
	// !이동=759364130569584640.954049652003602452
	if slice[0] == "!이동" {
		tmp := strings.Split(slice[1], ".")
		s.GuildMemberMove(m.GuildID, tmp[0], &tmp[1])
	}

	// !채널=954049652003602452
	if slice[0] == "!채널" {
		guild, _ := s.State.Guild(m.GuildID)
		voices := guild.VoiceStates

		print(len(voices))
		// n := 0
		// for n < len(voices) {
		// 	log(s, voices[n].UserID)
		// }
	}

	msg := fmt.Sprintf("%s %s %s %s", m.Timestamp.Format("20060102 15:04:05"), currntChannelName, m.Author.Username, m.Content)

	// 채팅 로그 채널에 메시지 전송
	s.ChannelMessageSend("952040735090294804", msg)
	// 출력
	fmt.Println(msg)
}
