use bbaddabot;

## 테이블 기본 조회
SELECT * FROM bbaddabot.channel;

SELECT * FROM bbaddabot.history;

SELECT * FROM bbaddabot.studyTotal;

SELECT * FROM bbaddabot.user;

## 오늘 데이터 조회
SELECT * FROM bbaddabot.studyTotal where date(date) = date(now());

SELECT * FROM bbaddabot.history where date(time) = date(now());

## 총 데이터 조회
SELECT COUNT(*) FROM bbaddabot.history;

## 개인별 일일 공부 시간 조회
SELECT * FROM bbaddabot.studyTotal
WHERE userNum = (SELECT userNum FROM bbaddabot.user WHERE userName="박현준" AND guildId = "951671348298661938");

## 길드 아이디, 유저 아이디로 오늘 (0 ~ 24) 공부한 시간 조회
select studytime from bbaddabot.studyTotal
WHERE date =DATE(NOW()) and 
usernum = (select usernum from user 
    where userid='759364130569584640' and guildid = '951671348298661938');

## 06시 기준으로 조회 하도록 변경 할 것

## 빠다 초기화
update user set bbadda=0;

## 주간 통계 / 원하는 값 X, 그래도 나중에 할때 도움될 것 같아 기록
## https://deeplify.dev/database/troubleshoot/mysql-daily-weekly-monthly
SELECT 
userNum,
DATE_FORMAT(DATE_SUB(`date`, INTERVAL (DAYOFWEEK(`date`)-1) DAY), '%Y/%m/%d') as start, 
DATE_FORMAT(DATE_SUB(`date`, INTERVAL (DAYOFWEEK(`date`)-7) DAY), '%Y/%m/%d') as end, 
DATE_FORMAT(`date`, '%Y%U') AS `date`, sum(`studyTime`)  as weekStudy
FROM bbaddabot.studyTotal GROUP BY userNum;

# 열 삭제
alter table bbaddabot.studytotal drop column todaySuccess;
alter table bbaddabot.user drop column wantTime;
alter table bbaddabot.user drop column  wantCnt;

# 열 추가
alter table bbaddabot.user add wantTime int default 180;
alter table bbaddabot.user add  wantCnt int default 5;

alter table bbaddabot.studytotal add todaySuccess boolean default false;
alter table bbaddabot.studytotal add weekSuccessCnt int default 0;

