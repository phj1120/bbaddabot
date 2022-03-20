use bbaddabot;

create table user(
	userNum int auto_increment,
	userId varchar(20),
    guildId varchar(20),
    userName varchar(50),
    bbadda int,
    userType varchar(20),
    
    primary key (userNum)
);

create table history(
	no int auto_increment,
    userNum int, 
	beforeChannelId varchar(20),
    afterChannelId varchar(20),
    time datetime not null,
    historyType varchar(20),
    
    primary key (`no`),
    foreign key (userNum) references user(userNum) ON DELETE CASCADE
);

create table studyTotal(
	no int auto_increment,
    userNum int,
    studyTime int not null,
    date datetime not null,
    
    primary key (no),
    foreign key (userNum) references user(userNum) ON DELETE CASCADE
);

create table channel(
	no int auto_increment,
    guildId varchar(20) not null,
    channelId varchar(20) not null,
    channelName varchar(20) not null,
    channelType varchar(20),
    
    primary key (no)
)