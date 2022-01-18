drop database if exists team;

create database team;

use team;

drop  table if exists user;
create table user(
    id int primary key auto_increment,
    nickname varchar(255) not null unique,
    password varchar(255) not null,
    avatar varchar(255) not null,
    feedback varchar(255) 
);

drop table if exists team;
create table team(
    id int primary key auto_increment,
    name varchar(255) not null unique,
    avatar varchar(255) not null,
    creator varchar(255) not null,
    team_coding varchar(255),
    foreign key (creator) references user(nickname)
);

drop table if exists user_team;
create table user_team(
    id int primary key auto_increment,
    username varchar(255) not null,
    teamname varchar(255) not null,
    foreign key (username) references user(nickname),
    foreign key (teamname) references team(name),
    unique (username,teamname)
);

drop table if exists project;
create table project(
    id int primary key auto_increment,
    name varchar(255) not null unique,
    creator varchar(255) not null,
    create_time varchar(255) not null,
    start_time varchar(255) not null,
    deadline varchar(255) not null,
    remark varchar(255) not null,
    team varchar(255) not null,
    foreign key(team) references team(name)
);

drop table if exists step;
create table step(
    id int primary key auto_increment,
    name varchar(255) not null unique,
    create_time varchar(255) not null,
    start_time varchar(255) not null,
    deadline varchar(255) not null,
    remark varchar(255) not null,
    project varchar(255) not null,
    foreign key (project) references project(name)
);

drop table if exists task;
create table task(
    id int primary key auto_increment,
    content varchar(255) not null,
    step varchar(255) not null,
    foreign key (step) references step(name)
);

drop table if exists user_task;
create table user_task(
    id int primary key auto_increment,
    principal varchar(255) not null,
    task_id int not null,
    performance varchar(5),
    foreign key (principal) references user(nickname),
    foreign key (task_id) references task(id)
);