select * from study_total;

insert into study_total(username, study_time, date) 
values ('hj', 2, DATE(now()));

update study_total 
set study_time = '1'
where username = 'hj';

INSERT INTO  user (userId, guildId, userName, bbadda, userType) 
	VALUES(' ', ' ', ' ', 0, 'user');
    
SELECT no, userNum, studyTime, date
			FROM studyTotal
			WHERE date=DATE(NOW()) AND userNum=1;
    
SELECT *
FROM user WHERE userId='759364130569584640' AND guildId='951671348298661938';

SELECT LAST_INSERT_ID();

SELECT TIMESTAMPDIFF(MINUTE,
		(SELECT time FROM history
		WHERE no = 25-1),
		(SELECT time FROM history
		WHERE no = 25
		));

SELECT * FROM history;

SELECT no, userNum, studyTime, date
FROM studyTotal
WHERE date=DATE(NOW()) AND userNum=1;

SELECT channelName
			FROM channel
			WHERE channelId = '954049652003602452';
            
SELECT studyTime
FROM studyTotal
WHERE date=DATE(NOW()) AND userNum=1;




-- UserNum 을 받아서 최신의 값 2개만 가져오고 그 값으로 비교해보자.
-- userNUm 으로 최신 값 2개 가져오기

-- 인덱스라 많아도 빨리 가져옴 
select *
from history use index for order by (HISTORYTIMEASC)
where userNum = 1;

select *
from history use index for order by (HISTORYTIMEDESC)
where userNum = 1;

select *
from history use index for order by (PRIMARY)
where userNum = 1;

-- order by 지양하라는데.... ㄱㅊ

SELECT *
FROM (SELECT time FROM history 
where userNum = 1 and date(time) = date(now()) 
ORDER BY time DESC LIMIT 2) AS rt;


SELECT TIMESTAMPDIFF(MINUTE,
		(SELECT time FROM recent
		WHERE 'rank' = 1),
		(SELECT time FROM recent
		WHERE 'rank' = 2))
FROM (SELECT time, rank() over (order by time desc) as 'rank'
		FROM history 
		where userNum = 1 and date(time) = date(now())) recent;




SELECT time, rank() over (order by time desc) as 'rank'
FROM history 
where userNum = 1 and date(time) = date(now());

SELECT TIMESTAMPDIFF(MINUTE,
(select time
from (SELECT time, rank() over (order by time desc) as 'rank'
		FROM history
		where userNum = 1 and date(time) = date(now())) as recent
where recent.rank =2),
(select time
from (SELECT time, rank() over (order by time desc) as 'rank'
		FROM history
		where userNum = 1 and date(time) = date(now())) as recent
where recent.rank =1)) as time;


select time
from (SELECT time, rank() over (order by time desc) as 'rank'
		FROM history
		where userNum = 1 and date(time) = date(now())) as recent
where recent.rank <=2;



create index HISTORYTIMEASC ON history (time ASC);
create index HISTORYTIMEDESC ON history (time DESC);

SELECT LAST_INSERT_ID() AS reviewIdx;


-- 시간 비교
SELECT TIMESTAMPDIFF(MINUTE,
		(SELECT *
		FROM (SELECT time FROM history 
		where userNum = 1 and date(time) = date(now()) 
		ORDER BY time DESC LIMIT 2) AS rt
));




