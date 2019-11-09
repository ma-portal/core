create table POSTING (
    ID char(36) primary key,
    CREATORID char(36) not null,
    CREATETIME int(64) not null,
    TYPE int(2) not null,
    CONTENT blob not null
);

create table COMMIT (
    ID char(36) primary key,
    CREATORID char(36) not null,
    CREATETIME int(64) not null,
    MAINREFID char(36) not null,
    REFID char(36) not null,
    CONTENT blob not null;
);

create table ACTIVATION (
    ID char(36) primary key,
    REFID char(36) not null,
    CREATETIME int(64) not null
);

create table TEACHER (
    ID char(36) primary key,
    ACCOUNT varchar(32) not null unique,
    PASSWORD varchar(32) not null
);

create table MEMBER (
    ID char(36) primary key,
    ACCOUNT varchar(32) not null unique,
    PASSWORD varchar(32) not null,
    ROLE int(2) not null
);