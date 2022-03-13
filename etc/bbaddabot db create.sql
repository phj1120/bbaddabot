create database bbaddabot;

create user 'bbadda'@'%' identified by 'mybbadda';
grant all on bbaddabot.* to 'bbadda'@'%';

show grants for 'bbadda'@'%';
