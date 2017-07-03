
create table tournaments
(
  id serial not null
    constraint toutnaments_pkey
    primary key,
  tournament_id varchar(64) not null,
  deposit bigint not null
)
;

create unique index toutnaments_tournament_id_uindex
  on tournaments (tournament_id)
;


create table join_events
(
  id bigserial not null
    constraint join_events_pkey
    primary key,
  tournament_id varchar(64) not null
    constraint join_events_tournaments_tournament_id_fk
    references tournaments (tournament_id)
    on update cascade on delete cascade,
  player_id varchar(64) not null,
  backers jsonb
)
;

create unique index join_events_tournament_id_player_id_uindex
  on join_events (tournament_id, player_id)
;





create table money_transactions
(
  id bigserial not null
    constraint money_transactions_pkey
    primary key,
  player_id varchar(64) not null,
  type smallint not null,
  sum bigint not null,
  balance bigint default 0 not null,
  last_tx boolean default false not null
)
;

create index money_transactions_player_id_last_tx_index
  on money_transactions (player_id, last_tx)
;

