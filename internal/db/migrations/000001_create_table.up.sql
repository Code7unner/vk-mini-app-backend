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

create table if not exists matches
(
    id               bigint  not null primary key,
    team_left_id     bigint references teams(id),
    team_right_id    bigint references teams(id),
    time_created     int     not null,
    time_started     int     not null,
    team_left_ready  boolean not null,
    team_right_ready boolean not null
);

create table if not exists steams
(
    id                         bigint not null primary key,
    community_visibility_state int       not null,
    profile_state              int       not null,
    persona_name               varchar   not null,
    comment_permission         int       not null,
    profile_url                varchar   not null,
    avatar                     varchar,
    avatar_medium              varchar,
    avatar_full                varchar,
    avatar_hash                varchar,
    last_logoff                int       not null,
    persona_state              int,
    real_name                  varchar,
    primary_clan_id            varchar,
    time_created               int       not null,
    persona_state_flags        int,
    loc_country_code           varchar
);