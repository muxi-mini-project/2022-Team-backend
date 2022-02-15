drop database if exists team;

create database team;

use team;

drop  table if exists user;
create table user(
    id int primary key auto_increment,
    student_id varchar(255),
    nickname varchar(255) not null,
    phone varchar(11) not null unique,
    password varchar(255) not null,
    feedback varchar(255),
    avatar varchar(255) null,
    sha   varchar(255) null,
    path varchar(255) null
);

drop table if exists team;
create table team(
    id int primary key auto_increment,
    name varchar(255) not null,
    creator_id int not null,
    team_coding varchar(255),
    avatar varchar(255) null,
    foreign key (creator_id) references user(id)
);

drop table if exists user_team;
create table user_team(
    id int primary key auto_increment,
    user_id int not null,
    team_id int not null,
    foreign key (user_id) references user(id),
    foreign key (team_id) references team(id),
    unique (user_id,team_id)
);

drop table if exists project;
create table project(
    id int primary key auto_increment,
    name varchar(255) not null unique,
    creator_id int not null,
    create_time varchar(255) not null,
    start_time varchar(255) not null,
    deadline varchar(255) not null,
    remark varchar(255) not null,
    team_id int not null,
    foreign key(team_id) references team(id)
);

drop table if exists step;
create table step(
    id int primary key auto_increment,
    name varchar(255) not null,
    project_id int not null,
    foreign key (project_id) references project(id)
);

drop table if exists task;
create table task(
    id int primary key auto_increment,
    name varchar(255) not null,
    creator_id int not null,
    create_time varchar(255) not null,
    start_time varchar(255) not null,
    deadline varchar(255) not null,
    remark varchar(255) not null,
    step_id int not null,
    project varchar(255),
    step  varchar(255),
    team_id int not null,
    foreign key (creator_id) references user(id),
    foreign key (step_id) references step(id)
);

drop table if exists user_task;
create table user_task(
    id int primary key auto_increment,
    principal_id int not null,
    principal_name varchar(255),
    task_id int not null,
    performance varchar(5),
    foreign key (principal_id) references user(id),
    foreign key (task_id) references task(id)
);