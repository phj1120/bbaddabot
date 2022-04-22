/*
작성자 : 박현준
작성일 : 2022.03.19.

수정자 : 박현준
수정일 : 2022.03.19.

파일 설명
목표 달성 못한 유저 강퇴
*/
package business

// import (
// 	"bbaddabot/persistence"
// 	"fmt"

// 	"github.com/bwmarrin/discordgo"
// )

// // TODO 빠따 삭제로 관련 기능 삭제. 추후 DB 에 맞게 변경 할 것.
// func ThrowOut(s *discordgo.Session, bbaddaLimit int) {
// 	guildList, _ := persistence.SelectGuildIdList()
// 	fmt.Println(guildList)
// 	for _, guildId := range guildList {
// 		userList, _ := persistence.SelectOverBbadaa(guildId, bbaddaLimit)
// 		fmt.Println(guildId)
// 		for _, user := range userList {
// 			fmt.Println(user)
// 			err := s.GuildMemberDeleteWithReason(guildId, user.UserId, "빠따 초과")
// 			if err != nil {
// 				fmt.Println("GuildMemberDeleteWithReason err")
// 				fmt.Println(err)
// 			} else {
// 				persistence.DeleteUser(user.UserId)
// 			}
// 		}
// 	}
// }
