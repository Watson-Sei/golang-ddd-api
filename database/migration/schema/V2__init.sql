create table users
(
    id int primary key,
    name text,
    created_at timestamp default (datetime(current_timestamp,'localtime'))
);