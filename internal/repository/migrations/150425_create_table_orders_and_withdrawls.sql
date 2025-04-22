-- +goose Up
-- +goose StatementBegin
CREATE
TYPE order_status_type AS ENUM ('UNKNOWN', 'REGISTRED', 'INVALID', 'PROCESSING', 'PROCESSED');

create table if not exists orders
(
    id          varchar(50) primary key,
    accrual     numeric(12, 2),
    status order_status_type default 'UNKNOWN',
    uploaded_at TIMESTAMP   not null,
    user_id     varchar(50) not null,
    UNIQUE (id, user_id)
);

create table if not exists withdrawals
(
    order_id     varchar(50) primary key,
    sum          numeric(12, 2) not null,
    processed_at TIMESTAMP      not null,
    user_id      varchar(50)    not null
);
-- +goose StatementEnd