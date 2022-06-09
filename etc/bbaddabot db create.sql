# 디비 생성
create database lsa_test;

# 사용자 생성
create user 'lsa'@'%' identified by 'lsa_pass';

# 사용자에게 DB의 모든 테이블을 all 제어할 수 있는 권한 부여
grant all on lsa_test.* to 'lsa'@'%';


show grants for 'lsa'@'%';
