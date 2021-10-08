
--  drop database if exists location_db;

create database if not exists location_db;

use location_db;

create table if not exists continents(	
	code varchar(2) not null,	
	name varchar(50) not null,
	description text null,
	primary key (code)
);

create table if not exists countries(
	code2 varchar(2) not null,
	code3 varchar(3) null,
	continent_code varchar(2) not null,
	name varchar(50) not null,
	description text null,
	primary key (code2),
	foreign key (continent_code) references continents(code)	
);

create table if not exists provinces(
	code varchar(2) not null,	
	name varchar(50) not null,
	country_code varchar(2) not null,	
	primary key (code),
	foreign key (country_code) references countries(code2)
);

