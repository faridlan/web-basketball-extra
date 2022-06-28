show tables;
create table users (
id int not null auto_increment primary key,
username varchar(100) not null,
email varchar(100) not null,
password varchar(225) not null,
role_id int not null,
status_id int not null default 1
)engine=innodb;

show tables;
desc users;

create table roles(
id int not null auto_increment primary key,
name varchar(50) not null
)engine=InnoDB;

create table status(
id int not null auto_increment primary key,
name varchar(50) not null
)engine=InnoDB;

create table biodata(
id int not null auto_increment primary key,
fullname varchar(200) not null,
age int not null,
year bigint not null,
position varchar(100) default '',
user_id int not null
)engine=InnoDB;

create table data_bio(
id int not null auto_increment primary key,
birth_certi varchar(255) not null,
family_card varchar(255) not null,
identity_card varchar(255) not null,
nisn varchar(255) not null,
profile_report varchar(255) not null,
image varchar(255) not null,
bio_id int not null
)engine=InnoDB;

create table registarion(
id int not null auto_increment primary key,
status enum('Belum', 'Menunggu', 'Selesai'),
user_id int not null
)engine=InnoDB;

create table dues (
id int not null auto_increment primary key,
presence enum('Hadir', 'Tidak'),
total_dues int not null default 0,
dues_date bigint not null,
user_id int not null
)engine=InnoDB;

create table income(
id int not null auto_increment primary key,
type enum('Pendaftaran', 'Iuran', 'Sumbangan', 'Hadiah'),
total int not null,
income_date bigint not null,
description varchar(200) not null default ''
)engine=InnoDB;

create table outcome(
id int not null auto_increment primary key,
description varchar(200) not null default '',
total int not null,
outcome_date bigint not null
)engine=InnoDB;

create table achievement(
id int not null auto_increment primary key,
description varchar(200) not null default ''
)engine=InnoDB;

