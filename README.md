## 디스코드 공부 시간 관리 봇

### 필요 모듈 다운로드
```
go mod tidy
```

### 실행
```
go run .
```

### 참고 자료
Go
* https://go.dev/doc/tutorial/getting-started

Discordgo
* https://pkg.go.dev/github.com/bwmarrin/discordgo#section-readme

Go layered Arcitecture
* https://github.com/ruslantsyganok/clean_arcitecture_golang_example


## 기록
https://parkhj.notion.site/da3f8fe5a214450386461b4e2f7f33e6
### 2022.03.16.
```
금방 끝날 줄 알았는데 점점 길어짐
요구사항 분석, 설계를 어느 정도 했다 생각했는데 부족한 부분들이 있었다.
일단 코드 부터 치지말고 
`비즈니스가 어떻게 진행되는지 말로 상세하게 적어보고 그 로직을 코드로 옮기기만 하면 될때 코딩 시작`하자.
이때 Usecase 를 쓰는 건가?
그래도 오늘 datasource와 persistence 를 어느정도 작성했기에 부족한 부분을 볼 수 있었음.
아니었으면 계속 해멨을듯
```

2022.03.19.
## 테이블 설계
user(userNum, userId, guildId, userName, bbadda, type)
	type{관리자, 사용자}

history(no, userNum, beforeChannelId, afterChannelId, time, type)
	type{입장, 퇴장, 공부, 휴식}

studyTotal(no, userNum, studyTime, date)

channel(no, guildId, channelId, channelName, type)
	type{공부, 휴식}


## 요구사항 분석
- 사용자 음성 채팅방 이동 → 사용자 상태 업데이트 됐다는 신호 발생
```        
        1. 이전 채널이 없는 경우 입장으로 판단
        
        이전 채널의 타입이 휴식인 경우 : 휴식 시작(기록 업데이트)
        
        이전 채널의 타입이 공부인 경우 : 공부 시작(기록 업데이트)
        
        2.이전 채널, 변경된 채널 모두 존재하면 채널 변경으로 판단
        
        이전 채널 타입이 공부
        
        이후 채널 타입이 공부 → 공부로 판단 : 공부 완료 (기록, 총 공부 시간 업데이트)
        
        이후 채널 타입이 휴식 → 공부로 판단 : 공부 완료, 총 공부 시간 업데이트
        
        이전 채널 타입이 휴식
        
        이후 채널 타입이 공부 → 휴식으로 판단 : 기록 업데이트 (휴식 완료)
        
        이후 채널 타입이 휴식 → 휴식으로 판단 : 기록 업데이트 (휴식 완료)
        
        3.이후 채널이 없는 경우 퇴장으로 판단
        
        이전 채널의 타입이 휴식인 경우 : 기록 업데이트(휴식 완료)
        
        이전 채널의 타입이 공부인 경우 :  기록 업데이트(공부 완료), 총 공부 시간 업데이트 
```
    변경이 발생할 때마다 기록
    
    이전 채널이 있는 경우 이전 채널의 타입으로 공부, 휴식 시간 구분
    
    DB 의 이전 row 에서 시간을 가져와 활동 시간 계산
    
    → 대략적인 정보 log 채널에 전송
    
    해당 활동이 공부 일 경우 총 공부 시간 업데이트


- 특정 시간에 하루 공부 할당량 확인
    
    → 빠따 계산(못 채운 사람 빠따 + 1 / 채운 사람 빠따 -1)
    
    → 빠따 정산(빠따가 n 초과 인 경우 강퇴 / 최소 빠따 0, 최대 빠따 n)
    

- 사용자 입력
    
    !공부시간 → 총 공부 시간 알림
    
    !빠따 → 현재 빠따 수 알림 
    
    !기록 → 현재 길드원들 빠따 수, 공부 시간 알림
    
- 관리자 입력
    
    !설정.휴식채널=채널ID → 해당 채널 휴식 채널로 지정
    
    !설정.공부채널=채널ID → 해당 채널 공부 채널로 지정
    
    !설정.관리자.지정=유저ID
    
    !설정.관리자.해제=유저ID


계층 구조
### Datastruct
- 각 DB 테이블 속성에 매칭되는 struct

- channel / history / studyTotal / user 

### persistence
- DB 에 접근해 CRUD 하는 계층

- dbconn / channelDao / histotyDao / studyTotalDao / userDao

### business
- 비즈니스 로직이 담겨있는 계층

- persistence 의 메서드 동작 여부 테스트

- ChangeChannel / test



### presentation
- 사용자로부터 데이터를 받거나 서비스를 제공하는 계층

- bbaddabot

## 수정 사항

채널 이동 시 상황에 맞게 상태 업데이트 하는 기능 buisness 계층으로 이동


# 2022.03.19
Spring 강의를 듣는데 

Repository - 순수 Data Access 기능

Service - 비즈니스 로직, 트랜잭션 처리

Controller - 요청과 응답 처리, 데이터 유효성 검증, 실행 흐름 제어

DTO - 계층간 데이터 공유

이렇게 구성 되어 있다해서 이에 맞게 구성해 보려 했다.

예외처리를 Controller 에 집중해서 하기 위해 다른 계층에서 예외처리를 안했다.

하지만 golang 은 자바와 다르게 try-catch 가 없어서 메서드에서 error 를 반환하고 `if error != nil { return } ` 이렇게 처리한다.

그래도 일단 예외처리를 모으자는 생각으로 예외처리를 없앴다.

하지만 거의 다 작성하고 보니 잘 못 되었음을 알게 되었다.

예외가 발생할만한 상황 마다 예외처리를 해줬어야한다.

수정하기에는 너무 많은 리소스가 들 것 같아 잘 못된 것을 알지만 안고가기로 했다.

사소하다고 생각해서 넘어간 개발자의 안일함과 무지가 이렇게 크게 돌아올 수 있다는 것을 알게되었다.


이 프로젝트를 하면서 지금까지는 간단한 CRUD 만 해봐서 Service 계층이 필요할까 라는 의문이 들었는데

직접 구성해보니 왜 나누는지를 어느 정도 알 수 있었다.

그리고 Repository 와 Service 에서 어떤 기준으로 나눠야하는지에 대한 궁금증이 생겼다.

쿼리로 가져오면 금방 가져올 것 같은데 

새로 DTO 와 Repository 를 만드는게 좋을까

기존에 있는 메서드들로 Service 에서 값을 처리하는게 좋을까 모르겠다.

좋은 코드들을 보면서 감을 익히고 싶다. 