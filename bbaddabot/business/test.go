package business

import (
	ds "bbaddabot/datastruct"
	"bbaddabot/persistence"
	"fmt"
	"time"
)

func UserDaoTest() {
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

}

func HistoryDaoTest() {
	fmt.Println("HistoryTest")
	history := ds.History{}
	history.UserNum = 1
	history.BeforeChannelId = "951672059010879499"
	history.AfterChannelId = "951672831312289883"
	history.HistoryType = "공부"
	cnt, _ := persistence.InsertHistory(history)
	fmt.Println(cnt)

	historys := persistence.SelectTodayHistoryByUserNum(history.UserNum)
	fmt.Println(historys)

	fmt.Println("SelectMinuteSpentChannel")
	spentMinute := persistence.SelectMinuteSpentByUserNum(cnt)
	fmt.Println(spentMinute)
}

func StudyTotalDaoTest() {
	fmt.Println("StudyTotalDaoTest")
	studyTotal := ds.StudyTotal{}
	studyTotal.UserNum = 1
	studyTotal.Study_time = 0
	studyTotal.Date = time.Now()

	fmt.Print("SelectStudyTotalTodayByUserNum")
	studyTime, err := persistence.SelectStudyTotalTodayByUserNum(studyTotal.UserNum)

	// 오늘 공부가 처음이면 오늘 기록 새로 생성 및 조회
	if err != nil {
		fmt.Println("InsertNewStudyTotal ")
		id := persistence.InsertNewStudyTotal(studyTotal.UserNum, studyTotal.Study_time)
		fmt.Println(id)
		studyTime, _ = persistence.SelectStudyTotalTodayByUserNum(studyTotal.UserNum)
	}
	fmt.Println(studyTime)

	// 공부 시간 업데이트
	persistence.UpdateStudyTimeByUserNumAndStudyTime(studyTotal.UserNum, 10)
	studyTime, _ = persistence.SelectStudyTotalTodayByUserNum(studyTotal.UserNum)
	fmt.Println(studyTime)
}

func ChannelTest() {
	channel := ds.Channel{}
	channel.GuildId = "951671348298661938"
	channel.ChannelId = "951672059010879499"
	channel.ChannelType = "study"

	fmt.Println("SelectChannelNameById")
	channelName := persistence.SelectChannelNameById(channel.ChannelId)

	if channelName == "" {
		fmt.Print("InsertChannel")
		persistence.InsertChannel(channel)
		persistence.SelectChannelNameById(channel.ChannelId)
	}
	fmt.Print(channelName)

	fmt.Println("UpdateChannelType")
	cnt := persistence.UpdateChannelType(channel.ChannelId, "rest")
	fmt.Println(cnt)
}
