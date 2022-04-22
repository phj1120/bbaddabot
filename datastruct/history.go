/*
작성자 : 박현준
작성일 : 2022.03.19.

수정자 : 박현준
작성일 : 2022.03.19.

파일 설명
History 타입 선언
*/

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
