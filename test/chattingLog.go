// https://pkg.go.dev/github.com/bwmarrin/discordgo#section-readme

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	// "time"
	// "reflect"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// 환경 변수에 토큰 추가. 처음 설정시 터미널 다시 열어도 안 되서 재부팅 하니까 됐음
	Token := os.Getenv("BBADDABOTTOKEN")
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	// dg.AddHandler(voiceConnection)
	// dg.AddHandler(ready)

	// Just like the ping pong example, we only care about receiving message
	// events in this example.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	currntChannel, err := s.Channel(m.ChannelID)
	currntChannelName := currntChannel.Name

	msg := fmt.Sprintf("%s %s %s %s", m.Timestamp.Format("20060102 15:04:05"), currntChannelName, m.Author.Username , m.Content)

	// 채팅 로그 채널에 메시지 전송
	_, err = s.ChannelMessageSend("952040735090294804", msg)
	// 출력
	fmt.Println(msg)

	// guildId := "951671348298661938"
	// 길드 정보 확인
	

	// guild, err := s.Guild(guildId)
	// fmt.Println(guild)
	
	// guildAduitLog, err := s.GuildAuditLog(guildId, "759364130569584640", "759364130569584640", 4, 50)
	// fmt.Println(guildAduitLog)
	
	if m.Content != "ping" {
		return
	}
	_, err = s.ChannelMessageSend(m.ChannelID, "Pong!")

	if err != nil {
		fmt.Println("error send Message:", err)
		return
	}
}