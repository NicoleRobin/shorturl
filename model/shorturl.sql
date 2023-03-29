create table t_shorturl (
    id bigint auto_increment,
    url varchar(512) not null default "",
    short_url varchar(100) not null default "",
    primary key (`id`),
    unique key uk_url(`url`),
    unique key uk_short_url(`short_url`)
) engine=innodb, default charset=utf8mb4;