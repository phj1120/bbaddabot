use bbaddabot;

create table history(
	`no` int auto_increment,
    `username` varchar(20) not null,
    `before_channel` varchar(20),
    `after_channel` varchar(20),
    `time` datetime not null,
    
    primary key (`no`)
);

create table study_total(
	`no` int auto_increment,
    `username` varchar(20),
    `study_time` time not null,
    `date` datetime not null,
    
    primary key (`no`)
);

create table channel_total(
	`no` int auto_increment,
    username varchar(20),
    bbadda int,
    
    primary key (`no`)
);