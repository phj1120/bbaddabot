package presentation

import (
	"bbaddabot/business"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func PresentationTest() {
	fmt.Println("PresentationTest")
	// business.UserDaoTest()
	// business.HistoryDaoTest()
	// business.StudyTotalDaoTest()
	business.ChannelTest()
	// business.ChangeChannel()
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
	dg.AddHandler(vociceStatusUpdate)
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

func vociceStatusUpdate(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	if v.VoiceState == nil {
		return
	}

	msg := business.ChangeChannel(s, *v)
	fmt.Println(msg)
	s.ChannelMessageSend("952057033476177920", msg)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// 강퇴 성공
	strs := m.Content
	// slice := strings.Split(strs, "=")

	if strs == "!test" {
		PresentationTest()
	}

	// if slice[0] == "!강퇴" {
	// 	s.GuildMemberDeleteWithReason(m.GuildID, slice[1], "강퇴")
	// }

	// // 제자리 이동 가능
	// // !이동=759364130569584640.954049652003602452
	// if slice[0] == "!이동" {
	// 	tmp := strings.Split(slice[1], ".")
	// 	s.GuildMemberMove(m.GuildID, tmp[0], &tmp[1])
	// }

	// // !채널=954049652003602452
	// if slice[0] == "!채널" {
	// 	guild, _ := s.State.Guild(m.GuildID)
	// 	voices := guild.VoiceStates

	// 	print(len(voices))
	// 	// n := 0
	// 	// for n < len(voices) {
	// 	// 	log(s, voices[n].UserID)
	// 	// }
	// }

	// 채팅 로그
	// currntChannel, _ := s.Channel(m.ChannelID)
	// currntChannelName := currntChannel.Name
	// msg := fmt.Sprintf("%s %s %s %s", m.Timestamp.Format("20060102 15:04:05"), currntChannelName, m.Author.Username, m.Content)
	// s.ChannelMessageSend("952040735090294804", msg)
	// fmt.Println(msg)
}

// 기타 기능
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
