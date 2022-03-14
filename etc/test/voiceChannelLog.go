package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	// "reflect"

	"github.com/bwmarrin/discordgo"
)

var token string
var buffer = make([][]byte, 0)

func main() {
	token := os.Getenv("BBADDABOTTOKEN")
	
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Register guildCreate as a callback for the guildCreate events.
	dg.AddHandler(guildCreate)

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


func myVociceStatusUpdate(s *discordgo.Session, v *discordgo.VoiceStateUpdate){

	// msg := fmt.Sprintf("%s%s",v.BeforeUpdate, v.VoiceState)
	// fmt.Println(msg)

	if v.VoiceState == nil{
		return
	}
	
	userName := getUserName(s, v.VoiceState.UserID)
	nowTime := time.Now().Format("20060102 15:04:05")
	msg := "None"

	if v.BeforeUpdate != nil && v.VoiceState.ChannelID == ""{
		msg = fmt.Sprintf("%s%s%s%s",nowTime," ", userName, " 종료")
		fmt.Println(msg)
	}else if v.BeforeUpdate != nil && v.VoiceState != nil{
		beforeChannelName := getChannelName(s, v.BeforeUpdate.ChannelID)
		afterChannelName := getChannelName(s, v.VoiceState.ChannelID)
		msg = fmt.Sprintf("%s%s%s%s%s%s%s", nowTime," ", userName ," 이동 : ",beforeChannelName, " -> ",afterChannelName)
		fmt.Println(msg)
	} else if v.BeforeUpdate == nil && v.VoiceState.ChannelID != ""{
		msg = fmt.Sprintf("%s%s%s%s",nowTime," ", userName, " 입장")
		fmt.Println(msg)
	}
	s.ChannelMessageSend("952057033476177920", msg)
}


func getChannelName(s *discordgo.Session, channelID string) string{
	channel, _ := s.Channel(channelID)
	return channel.Name
}

func getUserName(s *discordgo.Session, userID string) string{
	user, _ := s.User(userID)
	return user.Username
}


func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!공부시간") {
		// Find the channel that the message came from.
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			return
		}

		// Find the guild for that channel.
		g, err := s.State.Guild(c.GuildID)
		if err != nil {
			// Could not find guild.
			return
		}

		// Look for the message sender in that guild's current voice states.
		for _, vs := range g.VoiceStates {
			if vs.UserID == m.Author.ID {
				err = playSound(s, g.ID, vs.ChannelID)
				if err != nil {
					fmt.Println("Error playing sound:", err)
				}

				return
			}
		}
	}
}

// This function will be called (due to AddHandler above) every time a new
// guild is joined.
func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {

	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			fmt.Println(channel.ID)
			_, _ = s.ChannelMessageSend(channel.ID, "Airhorn is ready! Type !airhorn while in a voice channel to play a sound.")
			return
		}
	}
}


// playSound plays the current buffer to the provided channel.
func playSound(s *discordgo.Session, guildID, channelID string) (err error) {
	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}

	// Sleep for a specified amount of time before playing the sound
	time.Sleep(250 * time.Millisecond)

	// Start speaking.
	vc.Speaking(true)

	// Send the buffer data.
	for _, buff := range buffer {
		vc.OpusSend <- buff
	}

	// Stop speaking
	vc.Speaking(false)

	// Sleep for a specificed amount of time before ending.
	time.Sleep(250 * time.Millisecond)

	// Disconnect from the provided voice channel.
	vc.Disconnect()

	return nil
}