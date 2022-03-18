package datastruct

import (
	"time"
)

type History struct {
	No              int
	UserNum         int
	BeforeChannelId string
	AfterChannelId  string
	Time            time.Time
	HistoryType     string
}
