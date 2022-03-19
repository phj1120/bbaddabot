
######################## USER ########################
-- CREATE
INSERT INTO  user (userId, guildId, userName, bbadda, userType) 
VALUES('759364130569584640', '951671348298661938', '박현준', 0, 'overseer');

-- READ
SELECT userNum, userId, guildId, userName, bbadda, userType
FROM user
WHERE 1;

-- UPDATE
UPDATE user SET bbadda = bbadda+1history
WHERE userNum = 1;

UPDATE user SET userType = 'overseer'
WHERE userNum = 1;

-- DELETE 
DELETE FROM user
WHERE bbadda = 4;


######################## history ########################
-- CREATE
INSERT INTO  history (userNum, beforeChannelId, afterChannelId, time, historyType) 
VALUES(1, '951672059010879499', '951672831312289883', now(), 'study');

SET @LAST_HISTORY = (SELECT LAST_INSERT_ID());
SELECT @LAST_HISTORY;

-- READ
SELECT no, userNum, beforeChannelId, afterChannelId, time, historyType
FROM history
WHERE 1;



-- UPDATE
-- None

-- DELETE 
-- None

######################## studyTotal ########################
-- CREATE
INSERT INTO  studyTotal (userNum, studyTime, date) 
VALUES(1, 10, DATE(now()));

-- READ
SELECT no, userNum, studyTime, date
FROM studyTotal
WHERE date=DATE(NOW()) AND userNum=1;

-- UPDATE
UPDATE studyTotal 
SET studytime = studytime + (SELECT TIMESTAMPDIFF(MINUTE,
	(SELECT time FROM history
	WHERE no = (SELECT @LAST_HISTORY -1 )),
	(SELECT time FROM history
	WHERE no = (SELECT @LAST_HISTORY)
	)))
WHERE date=DATE(NOW()) AND userNum=1;

-- study time
SELECT TIMESTAMPDIFF(MINUTE,
	(SELECT time FROM history
	WHERE no = (SELECT @LAST_HISTORY -1 )),
	(SELECT time FROM history
	WHERE no = (SELECT @LAST_HISTORY)
	));

-- DELETE 
-- None


######################## channel ########################
-- CREATE
INSERT INTO  channel (guildId, channelId, channelName, channelType) 
VALUES('951671348298661938', '954047912420204664', 'test-공부', 'study');

INSERT INTO  channel (guildId, channelId, channelName, channelType) 
VALUES('951671348298661938', '954049652003602452', 'test-휴식', 'rest');

-- READ
SELECT *
FROM channel
WHERE channelId = '951672059010879499';

-- UPDATE
UPDATE channel SET channelType = 'study'
WHERE channelId = '951672059010879499';

-- DELETE 
DELETE FROM channel
WHERE channelId = '952034251174473748';


