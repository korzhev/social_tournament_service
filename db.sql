create table join_events
(
  id bigserial not null
    constraint join_events_pkey
    primary key,
  tournament_id bigint not null
    constraint join_events_tournaments_tournament_id_fk
    references tournaments (tournament_id)
    on update cascade on delete cascade,
  player_id bigint not null,
  backers bigint[]
)
;

create unique index join_events_player_id_uindex
  on join_events (player_id)
;



create table money_transactions
(
  id bigserial not null
    constraint money_transactions_pkey
    primary key,
  player_id bigint not null,
  type smallint not null,
  sum bigint not null
)
;

create index money_transactions_player_id_index
  on money_transactions (player_id)
;



create table tournaments
(
  id serial not null
    constraint toutnaments_pkey
    primary key,
  tournament_id bigint not null,
  deposit bigint not null
)
;

create unique index toutnaments_tournament_id_uindex
  on tournaments (tournament_id)
;

