create table if not exists teams
(
    id           int     not null primary key,
    title        varchar not null,
    tag          varchar not null,
    photo_small  varchar not null,
    photo_medium varchar not null,
    photo_big    varchar not null,
    rating       int     not null,
    match_id     int references matches (id)
);