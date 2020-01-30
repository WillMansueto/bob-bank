CREATE DATABASE blockcoin;

\c blockcoin;

CREATE TABLE IF NOT EXISTS users(
    uid bigserial primary key,
    nickname varchar(15) not null unique,
    email varchar(40) not null unique,
    password varchar(100) not null,
    status char(1) default '0',
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

CREATE TABLE IF NOT EXISTS WALLETS(
    public_key varchar(32) primary key unique not null,
    usr bigint not null,
    balance real default 0.0,
    updated_at timestamp default current_timestamp,
    constraint wakkets_usr_fk foreign key(usr)
    references users(uid)
    on delete cascade
    on update cascade
);

CREATE TABLE IF NOT EXISTS transactions(
    uid bigserial primary key,
    origin varchar(32) not null,
    target varchar(32) not null,
    cash real not null,
    message varchar(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    constraint transactions_orign_fk foreign key(origin)
    references WALLETS(public_key)
    on delete cascade
    on update cascade,
    constraint transactions_target_fk foreign key(target)
    references WALLETS(public_key)
    on delete cascade
    on update cascade
);