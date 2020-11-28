create table if not exists users
(
    id           bigint     not null primary key,
    name         varchar not null,
    lastname     varchar not null,
    city         varchar,
    country      varchar,
    sex          int     not null,
    timezone     int     not null,
    photo_small  varchar not null,
    photo_medium varchar not null,
    photo_big    varchar not null,
    team_id      bigint references teams(id),
    steam_id     bigint references steams(id)
);