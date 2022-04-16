package datastruct

import (
	"time"
)

type StudyTotal struct {
	No             int
	UserNum        int
	StudyTime      int
	Date           time.Time
	TodaySuccess   bool
	WeekSuccessCnt int
}
