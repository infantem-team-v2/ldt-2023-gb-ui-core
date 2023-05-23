create schema service;

create table service.auth
(
    id          serial
        primary key,
    name        varchar(54)  not null
        unique,
    public_key  varchar(256) not null
        unique,
    private_key varchar(512) not null
        unique,
    url         varchar(256)
);

create schema ui;

create table types
(
    id               serial
        primary key,
    name             varchar(54)           not null
        unique,
    comment          text,
    multiple_options boolean default false not null
);