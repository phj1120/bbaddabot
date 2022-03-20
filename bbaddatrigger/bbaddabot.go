package main

import (
	"bbaddabot/persistence"
	"fmt"
	"os"

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

	guildList, _ := persistence.SelectGuildIdList()
	fmt.Println(guildList)
	for _, guildId := range guildList {
		userList, _ := persistence.SelectOverBbadaa(guildId, 4)
		fmt.Println(guildId)
		for _, user := range userList {
			fmt.Println(user)
			err := s.GuildMemberDeleteWithReason(guildId, user.UserId, "빠따 초과")

			msg := fmt.Sprint("빠따 초과 : %+v", user)
			s.ChannelMessageSend("954051159797137520", msg)

			if err != nil {
				fmt.Println("GuildMemberDeleteWithReason err")
				fmt.Println(err)
			} else {
				persistence.DeleteUser(user.UserId)
			}
		}
	}

	offSignal = true
}
