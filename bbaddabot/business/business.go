package business

import (
	ds "bbaddabot/datastruct"
	"bbaddabot/persistence"
	"fmt"
)

func BusinessUserTest() {
	fmt.Println("BusinessTest")

	user := ds.User{}
	user.UserId = "759364130569584640"
	user.GuildId = "951671348298661938"
	user.UserName = "박현준"
	user.Bbadda = 0
	user.UserType = "overseer"

	persistence.InsertUser(user)

	userNum, _ := persistence.SelectUserNumByUserIdAndGuildId(user.UserId, user.GuildId)
	fmt.Println(userNum)

	user2, _ := persistence.SelectUserByUserNum(userNum)
	fmt.Println(user2)

	persistence.UpdateUserBbaddaByUserId(user.UserId)
}
