package business

import (
	"bbaddabot/persistence"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ThrowOut(s *discordgo.Session, bbaddaLimit int) bool {
	guildList, _ := persistence.SelectGuildIdList()
	fmt.Println(guildList)
	for _, guildId := range guildList {
		userList, _ := persistence.SelectOverBbadaa(guildId, bbaddaLimit)
		fmt.Println(guildId)
		for _, user := range userList {
			fmt.Println(user)
			err := s.GuildMemberDeleteWithReason(guildId, user.UserId, "빠따 초과")
			if err != nil {
				fmt.Println("GuildMemberDeleteWithReason err")
				fmt.Println(err)
			} else {
				persistence.DeleteUser(user.UserId)
			}
		}
	}
	return true
}
