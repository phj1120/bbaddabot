use bbaddabot;

INSERT INTO  user (userId, guildId, userName, bbadda, userType) 
VALUES('759364130569584640', '951671348298661938', '박현준', 0, 'overseer'),
('358480813681016832', '951671348298661938', '유민상', 0, 'overseer');

INSERT INTO  history (userNum, beforeChannelId, afterChannelId, time, historyType) 
VALUES(1, '951672059010879499', '951672831312289883', now(), 'study');

INSERT INTO  studyTotal (userNum, studyTime, date) 
VALUES(1, 10, DATE(now()));

INSERT INTO  channel (guildId, channelId, channelName, channelType) 
VALUES('951671348298661938', '955078188806058034', '개인', '공부'),
('951671348298661938', '955078258293104680', '프로젝트', '공부'),
('951671348298661938', '955078488329695252', '질의응답', '공부'),
('951671348298661938', '955078423762567218', '휴식', '휴식'),
('951671348298661938', '955079430773043281', '알림', '알림'),
('951671348298661938', '952057033476177920', 'study-log', '로그');

