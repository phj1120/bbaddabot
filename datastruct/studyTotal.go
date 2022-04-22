/*
작성자 : 박현준
작성일 : 2022.03.19.

수정자 : 박현준
작성일 : 2022.03.19.

파일 설명
StudyTotal 타입 선언
*/

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
