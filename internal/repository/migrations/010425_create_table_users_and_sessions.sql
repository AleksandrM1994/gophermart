-- +goose Up
-- +goose StatementBegin
create table if not exists users
(
    id       varchar(50) primary key,
    login    varchar(50) not null unique,
    password text        not null,
    balance  numeric(12, 2) default 0
);

CREATE INDEX users_on_login_password_idx ON users (login, password);

create table if not exists sessions
(
    cookie        text primary key,
    cookie_finish TIMESTAMP   not null,
    user_id       varchar(50) not null REFERENCES users(id)
);
-- +goose StatementEnd