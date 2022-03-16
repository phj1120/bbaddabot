package datastruct

import (
	"time"
)

type History struct {
	No             int       `db:"no"`
	Username       string    `db:"username"`
	Before_channel string    `db:"before_channel"`
	After_channel  string    `db:"after_channel"`
	Time           time.Time `db:"time"`
}
