create table if not exists matches
(
    id               int not null primary key,
    team_left_id     int       not null,
    team_right_id    int       not null,
    time_created     int       not null,
    time_started     int       not null,
    team_left_ready  boolean   not null,
    team_right_ready boolean   not null
);