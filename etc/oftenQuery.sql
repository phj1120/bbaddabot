use bbaddabot;

# 단순 조회
## 테이블 기본 조회
SELECT * FROM bbaddabot.channel;

SELECT * FROM bbaddabot.history;

SELECT * FROM bbaddabot.studyTotal;

SELECT * FROM bbaddabot.user;

## 오늘 데이터 조회
SELECT * FROM bbaddabot.studyTotal where date(date) = date(now());

SELECT * FROM bbaddabot.history where date(time) = date(now()) and userNum = 13;

## 총 데이터 조회
SELECT COUNT(*) FROM bbaddabot.history;


# Study Total 조회
## 개인별 일일 공부 시간 조회
SELECT * FROM bbaddabot.studyTotal
WHERE userNum = (SELECT userNum FROM bbaddabot.user WHERE userName="박현준" AND guildId = "951671348298661938")
AND date='2022-04-12';

SELECT * FROM bbaddabot.studyTotal
WHERE userNum = (SELECT userNum FROM bbaddabot.user WHERE userName="유민상" AND guildId = "951671348298661938")
AND date='2022-04-11';

## 주간 공부 현황 
SELECT
	date_format(date,'%Y-%m-%d') AS 일자, userName AS 이름, studyTime AS 공부시간
FROM
    studyTotal LEFT JOIN user ON studyTotal.userNum = user.userNum
WHERE
    date_format(date,'%Y-%m-%d')
    BETWEEN
        (SELECT ADDDATE(CURDATE(), - WEEKDAY(CURDATE()) + 0 ))
    AND
        (SELECT ADDDATE(CURDATE(), - WEEKDAY(CURDATE()) + 6 ));

## 이번 달
SELECT
	date_format(date,'%Y-%m-%d') AS 일자, userName AS 이름, studyTime AS 공부시간
FROM
    studyTotal LEFT JOIN user ON studyTotal.userNum = user.userNum
WHERE
    date_format(date,'%Y-%m-%d')
    BETWEEN
        (SELECT DATE_FORMAT(NOW(), '%x-%m-01'))
    AND
        (SELECT LAST_DAY(NOW()));

# 저번 달 ( 하드코딩 ) ------------------------------------------------------------------------------------------ #
SELECT
	date_format(date,'%Y-%m-%d') AS 일자, userName AS 이름, studyTime AS 공부시간
FROM
    studyTotal LEFT JOIN user ON studyTotal.userNum = user.userNum
WHERE
    date_format(date,'%Y-%m-%d')
    BETWEEN
        (SELECT DATE_FORMAT(NOW(), '2022-04-01'))
    AND
        (SELECT DATE_FORMAT(NOW(), '2022-04-30'));
        
SELECT
	userName, sum(studyTime) as 공부시간
FROM
    studyTotal LEFT JOIN user ON studyTotal.userNum = user.userNum
WHERE
    date_format(date,'%Y-%m-%d')
    BETWEEN
        (SELECT DATE_FORMAT(NOW(), '2022-04-01'))
    AND
        (SELECT DATE_FORMAT(NOW(), '2022-04-30'))
group by userName;
# ----------------------------------------------------------------------------------------------------------- #

## 길드 아이디, 유저 아이디로 오늘 (0 ~ 24) 공부한 시간 조회
select studytime from bbaddabot.studyTotal
WHERE date =DATE(NOW()) and 
usernum = (select usernum from user 
    where userid='759364130569584640' and guildid = '951671348298661938');

## 유저별 이번주 공부 시간
SELECT
    userNum, sum(studyTime)
FROM
    studyTotal
WHERE
    date_format(date,'%Y-%m-%d')
    BETWEEN
        (SELECT ADDDATE(CURDATE(), - WEEKDAY(CURDATE()) + 0 ))
    AND
        (SELECT ADDDATE(CURDATE(), - WEEKDAY(CURDATE()) + 6 ))
	AND
		userNum = '1';
       
## 유저별 이번달 공부 시간
SELECT
    sum(studyTime)
FROM
    studyTotal
WHERE
    date_format(date,'%Y-%m-%d')
    BETWEEN
        (SELECT DATE_FORMAT(NOW(), '%x-%m-01'))
    AND
        (SELECT LAST_DAY(NOW()))
	AND userNum = '1';


# DB 운영
select version();
show status;
show status like '%thread%';
