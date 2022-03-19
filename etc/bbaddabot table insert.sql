
INSERT INTO  user (userId, guildId, userName, bbadda, userType) 
VALUES('759364130569584640', '951671348298661938', '박현준', 0, 'overseer'),
('358480813681016832', '951671348298661938', '유민상', 0, 'overseer');

INSERT INTO  history (userNum, beforeChannelId, afterChannelId, time, historyType) 
VALUES(1, '951672059010879499', '951672831312289883', now(), 'study');

INSERT INTO  studyTotal (userNum, studyTime, date) 
VALUES(1, 10, DATE(now()));

INSERT INTO  channel (guildId, channelId, channelName, channelType) 
VALUES('951671348298661938', '951672059010879499', '공부', 'study'),
('951671348298661938', '951672831312289883', '휴식', 'rest'),
('951671348298661938', '951671548874489956', '프로젝트', 'study'),
('951671348298661938', '951673410822484088', '질의응답', 'study'),
('951671348298661938', '954047912420204664', 'test-공부', 'study'),
('951671348298661938', '954049652003602452', 'test-휴식', 'rest');



