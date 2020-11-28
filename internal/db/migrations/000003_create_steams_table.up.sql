create table if not exists steams
(
    id                         bigserial not null primary key,
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