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