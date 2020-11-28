create table if not exists public.users
(
    id           bigint  not null primary key,
    name         varchar not null,
    lastname     varchar not null,
    city         varchar,
    country      varchar,
    sex          int     not null,
    timezone     int     not null,
    photo_small  varchar not null,
    photo_medium varchar not null,
    photo_big    varchar not null,
    team_id      bigint,
    steam_id     bigint
);

create table if not exists public.teams
(
    id           bigint  not null primary key,
    title        varchar not null,
    tag          varchar not null,
    photo_small  varchar not null,
    photo_medium varchar not null,
    photo_big    varchar not null,
    rating       int     not null,
    match_id     bigint
);

create table if not exists public.matches
(
    id               bigint  not null primary key,
    team_left_id     bigint,
    team_right_id    bigint,
    time_created     int     not null,
    time_started     int     not null,
    team_left_ready  boolean not null,
    team_right_ready boolean not null
);

create table if not exists public.steams
(
    id                         bigint  not null primary key,
    community_visibility_state int     not null,
    profile_state              int     not null,
    persona_name               varchar not null,
    comment_permission         int     not null,
    profile_url                varchar not null,
    avatar                     varchar,
    avatar_medium              varchar,
    avatar_full                varchar,
    avatar_hash                varchar,
    last_logoff                int     not null,
    persona_state              int,
    real_name                  varchar,
    primary_clan_id            varchar,
    time_created               int     not null,
    persona_state_flags        int,
    loc_country_code           varchar
);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_team_id FOREIGN KEY (team_id) REFERENCES public.teams (id);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_steam_id FOREIGN KEY (steam_id) REFERENCES public.steams (id);

ALTER TABLE ONLY public.teams
    ADD CONSTRAINT fk_match_id FOREIGN KEY (match_id) REFERENCES public.matches (id);

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT fk_team_left_id FOREIGN KEY (team_left_id) REFERENCES public.teams (id);

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT fk_team_right_id FOREIGN KEY (team_right_id) REFERENCES public.teams (id);