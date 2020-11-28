create table if not exists teams
(
    id           bigint     not null primary key,
    title        varchar not null,
    tag          varchar not null,
    photo_small  varchar not null,
    photo_medium varchar not null,
    photo_big    varchar not null,
    rating       int     not null,
    match_id     bigint references matches (id)
);